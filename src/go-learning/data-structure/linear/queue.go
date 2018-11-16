package linear

import (
	"container/list"
	"strconv"
)

type Queue struct {
	common
}

func NewQueue(autoMutexLock bool) (l *Queue) {
	l = &Queue{}
	l.data = list.New()
	l.autoLock = autoMutexLock
	return
}

func (s *Queue)Enqueue(item interface{})(ok bool)  {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	s.data.PushBack(item)
	return true
}
func (s *Queue) Dequeue()(item interface{}, ok bool)  {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	first := s.data.Front()
	if first == nil {
		return nil, false
	}
	s.data.Remove(first)
	return first.Value, true
}
func (s *Queue) Peek() (item interface{}, ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	if s.IsEmpty() {
		return nil, false
	}
	first := s.data.Front()
	return first.Value, true
}

func (s *Queue) String() string {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return "Queue [" + strconv.Itoa(s.data.Len()) + "]"
}