package server

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var EOF = errors.New("EOF")

// 空结构体
var Exists = struct{}{}
var ServerMemo *memo = newMemo()
var started = false

type memo struct {
	data map[interface{}]struct{}
	mu   *sync.RWMutex
}

func newMemo(items ...interface{}) *memo {
	m := &memo{data: make(map[interface{}]struct{}), mu: &sync.RWMutex{}}
	m.add(items...)
	return m
}

func (m *memo) add(items ...interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, item := range items {
		m.data[item] = Exists
	}
	return nil
}

func (m *memo) contains(item interface{}) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.data[item]
	return ok
}

func (m *memo) size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}

func (m *memo) clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = make(map[interface{}]struct{})
}

// 定义路由函数
func router() {
	http.HandleFunc("/exists", func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("客户端异常")
		}
		data := strings.Split(string(body), ",") // "1, 2, 3, 4, 5"
		var res []string
		for _, item := range data {
			if !ServerMemo.contains(item) {
				ServerMemo.add(item)
				res = append(res, "false")
			} else {
				res = append(res, "true")
			}
		}
		rw.Write([]byte(strings.Join(res, ",")))
	})

	http.HandleFunc("/clear", func(rw http.ResponseWriter, r *http.Request) {
		ServerMemo.clear()
	})
}

func Server() {
	// 避免重复创建服务对象
	if started {
		return
	}

	started = true
	router() // 装载路由
	err := http.ListenAndServeTLS(":8080", "./ssl/cert.pem", "./ssl/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
