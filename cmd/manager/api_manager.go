package manager

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
	"vscale-task/cmd/storage"

	"vscale-task/cmd/providers"
)

type APIManager struct {
	client  providers.Client
	Storage storage.Storage
}

func NewAPIManager(cl providers.Client, storage storage.Storage) APIManager {
	return APIManager{
		client:  cl,
		Storage: storage,
	}
}

func (m *APIManager) CreateServerGroup(chAccepted chan<- int64, servReq *providers.CreateServerRequest, number int64) (err error) {
	if servReq == nil {
		return errors.New("nil pointer error")
	}
	if number <= 0 {
		return errors.New("number of server error")
	}
	//TODO Валидация параметров CreateServerRequest

	var (
		sleep     bool
		interrupt bool
		counter   int64
		wg        sync.WaitGroup
		groupID   = m.Storage.NextGroupID()
	)

	m.Storage.SetGroupStatus(groupID, storage.StatusAccepted)
	chAccepted <- groupID
	for counter < number && !interrupt {
		if sleep {
			time.Sleep(30 * time.Second)
			sleep = false
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			var servResp providers.CreateServerResponse
			servResp, err = m.client.CreateServer(servReq)

			if err != nil {
				switch err {
				case providers.ErrTooManyRequests:
					sleep = true
					return
				default:
					interrupt = true
					return
				}
			}
			m.Storage.AddServer(groupID, servResp.CTID)
			atomic.AddInt64(&counter, 1)
		}()

	}
	wg.Wait()

	if interrupt {
		m.Storage.SetGroupStatus(groupID, storage.StatusFailed)
		if err = m.DeleteServerGroup(groupID); err != nil {
			return ErrFatalApiError
		}
		return ErrFatalApiError
	}

	m.Storage.SetGroupStatus(groupID, storage.StatusComplete)
	return nil
}

func (m *APIManager) DeleteServerGroup(groupID int64) (err error) {
	ctidList, ok := m.Storage.GetServerList(groupID)
	if !ok {
		return ErrGroupIDNotFound
	}
	var (
		sleep     bool
		interrupt bool
		wg        sync.WaitGroup
	)

	for i := range ctidList {
		var ctid = ctidList[i]
		if interrupt {
			break
		}
		if sleep {
			time.Sleep(30 * time.Second)
			sleep = false
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			var servResp providers.DeleteServerResponse
			servResp, err = m.client.DeleteServer(ctid)

			if err != nil {
				switch err {
				case providers.ErrTooManyRequests:
					sleep = true
					return
				default:
					interrupt = false
					return
				}
			}

			m.Storage.RemoveServer(groupID, servResp.CTID)
		}()
	}

	if interrupt {
		return ErrDeleteServer
	}

	m.Storage.SetGroupStatus(groupID, storage.StatusDeleted)
	return nil
}
