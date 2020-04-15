package storage

import (
	"sync"
)

type MapStorage struct {
	*sync.RWMutex
	data map[int64][]int64
}

func NewStorage() *MapStorage {
	return &MapStorage{
		RWMutex: new(sync.RWMutex),
		data:    make(map[int64][]int64),
	}
}

func (s *MapStorage) AddServer(groupID, ctid int64) (ok bool) {
	s.Lock()
	defer s.Unlock()

	ctidList, ok := s.data[groupID]
	if !ok {
		ctidList = []int64{}
	}
	ctidList = append(ctidList, ctid)
	s.data[groupID] = ctidList
	return true
}

func (s *MapStorage) RemoveServer(groupID, ctid int64) (ok bool) {
	s.Lock()
	defer s.Unlock()

	ctidList, ok := s.data[groupID]
	if !ok {
		return false
	}

	for i := range ctidList {
		if ctidList[i] == ctid {
			ctidList = append(ctidList[:i], ctidList[i+1:]...)
			break
		}
	}

	s.data[groupID] = ctidList
	return true
}

func (s *MapStorage) GetServerList(groupID int64) ([]int64, bool) {
	s.RLock()
	defer s.RUnlock()

	ctidList, ok := s.data[groupID]
	if !ok {
		return nil, false
	}
	return ctidList, true
}

func (s *MapStorage) NextGroupID() int64 {
	s.Lock()
	defer s.Unlock()

	var maxKey int64
	for k, _ := range s.data {
		if k > maxKey {
			maxKey = k
		}
	}
	return maxKey + 1
}

func (s *MapStorage) RemoveGroupID(groupID int64) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, groupID)
}

func (s *MapStorage) GetGroupStatus(groupID int64) (string, bool) {
	return "", false
}

func (s *MapStorage) SetGroupStatus(groupID int64, status string) bool {
	return false
}
