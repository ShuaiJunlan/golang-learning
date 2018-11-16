package linear

import (
	"container/list"
	"sync"
)

type ListCommon interface {
	Len() int
	HasElement() bool
	IsEmpty() bool
}

type common struct {
	data     *list.List
	autoLock bool
	sync.Mutex
}

func (s *common) Len() int {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len()
}

func (s * common)IsEmpty() bool {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() == 0
}

func (s *common) HasElement() bool {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() > 0
}