package manager

import (
	"errors"
	"swap-free/src/golang.org/x/tools/go/analysis/passes/copylock/testdata/src/a"
	"sync"
	"time"
	"vscale-task/cmd/storage"

	"vscale-task/cmd/providers"
)

type APIManager struct {
	Client  providers.Client
	Storage storage.Storage
}

func NewAPIManager(cl *providers.Client, storage *storage.Storage) APIManager {
	return APIManager{
		Client:  *cl,
		Storage: *storage,
	}
}

func (m *APIManager) CreateServerGroup(servReq *providers.CreateServerRequest, number int) (groupID int64, err error) {
	if servReq == nil {
		return 0, errors.New("nil pointer error")
	}
	if number <= 0 {
		return 0, errors.New("number of server error")
	}

	var (
		sleep     = make(chan struct{}, 1)
		rollback  = make(chan struct{}, 1)
		interrupt bool
		wg        sync.WaitGroup
	)

	groupID   = m.Storage.NextGroupID()

	for counter := 0; counter < number && !interrupt; counter++ {
		select {
		case <-sleep:
			time.Sleep(30 * time.Second)
		case <-rollback:
			interrupt = true
		default:
			wg.Add(1)
			go func() {
				defer wg.Done()
				var servResp providers.CreateServerResponse
				servResp, err = m.Client.CreateServer(servReq)

				if err != nil {
					switch err {
					case providers.ErrTooManyRequests:
						sleep <- struct{}{}
						return
					default:
						rollback <- struct{}{}
						return
					}
				}

				m.Storage.AddServer(groupID, servResp.CTID)
			}()
		}
	}
	wg.Wait()
	if interrupt {
		return 0, ErrNeedRollback
	}

	return groupID, nil
}


func (m *APIManager) DeleteServerGroup(groupID int64) (err error) {
	ctidList, ok := m.Storage.GetServerList(groupID)
	if !ok {
		return ErrGroupIDNotFound
	}
	var (
		sleep     = make(chan struct{}, 1)
		errorOccurred  = make(chan struct{}, 1)
		wg        sync.WaitGroup
	)

	for _, ctid := range ctidList {
		select {
		case <-sleep:
			time.Sleep(30 * time.Second)
		case <-errorOccurred:
			return ErrDeleteServer
		default:
			wg.Add(1)
			go func() {
				defer wg.Done()
				var servResp providers.DeleteServerResponse
				servResp, err = m.Client.DeleteServer(ctid)

				if err != nil {
					switch err {
					case providers.ErrTooManyRequests:
						sleep <- struct{}{}
						return
					default:
						errorOccurred <- struct{}{}
						return
					}
				}

				m.Storage.AddServer(groupID, servResp.CTID)
			}()
		}
	}

	return nil
}