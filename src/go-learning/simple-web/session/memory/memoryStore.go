package memory

import (
	"container/list"
	"sync"
	"time"
)
var pder = &Provider{list:list.New()}
type SessionStore struct {
	sid string
	timeAccessed time.Time
	value map[interface{}] interface{}
}


func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	//pder.sessions
}
type Provider struct {
	lock sync.Mutex
	sessions map[string]*list.Element
	list *list.List
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	//sess
}
