package concurrent

import (
	"encoding/json"
	"sync"
	"testing"
	"time"
)

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

func TestCondBroadCast(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	go CondRead("reader1", cond)
	go CondRead("reader2", cond)
	go CondRead("reader3", cond)
	CondWrite("writer", cond)

	time.Sleep(time.Second * 3)
}
