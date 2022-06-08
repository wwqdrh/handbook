package concurrent

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

////////////////////
// sync.pool
////////////////////

// 用于复用对象

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

////////////////////
// cond
////////////////////

// 用于通知多个对象，(如果使用通道或者全局变量需要额外的复杂度)

var done = false

func CondRead(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func CondWrite(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}
