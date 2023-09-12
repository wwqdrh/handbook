package mapping

import (
	"log"
	"sync"
	"time"
)

// 实现阻塞读并发安全的map
// key不存在的话就get操作等待，直到key存在或者超时，保证并发安全
// 阻塞协程考虑channel 并发安全使用锁

type sp interface {
	Out(key string, val interface{})
	Rd(key string, timeout time.Duration) interface{}
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func NewMap() sp {
	return &Map{
		c:   make(map[string]*entry),
		rmx: &sync.RWMutex{},
	}
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.RLock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.RUnlock()
		return e.value
	} else if !ok {
		m.rmx.RUnlock()
		m.rmx.Lock()
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.rmx.RUnlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	}
}
