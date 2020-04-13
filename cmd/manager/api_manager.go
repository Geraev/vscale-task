package manager

import (
	"errors"
	"golang.org/x/net/html/atom"
	"sync"
	"time"
	"vscale-task/cmd/providers"
)

type APIManager struct {
	Client providers.Client
}

func NewAPIManager(cl *providers.Client) APIManager {
	return APIManager{
		Client: *cl,
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
		sleep    = make(chan struct{}, 1)
		rollback = make(chan struct{}, 1)
		interrupt bool
		wg sync.WaitGroup
	)

	for counter := 0; counter < number && !interrupt; counter++ {
		select {
		case <-sleep:
			time.Sleep(30 * time.Second)
		case <-rollback:
			wg.Wait()
			//TODO Удалить все серверы группы
			//TODO Возможно это надо делать не здесь. Это не его функция
			//TODO Здесь можно просто выйти с ошибкой
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
				//TODO Сохранять CTID с привязкой к groupID в хранилище
			}()
		}
	}

	return
}
