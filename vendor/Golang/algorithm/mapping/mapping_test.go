package mapping

import (
	"log"
	"testing"
	"time"
)

func TestConcurrentMap(t *testing.T) {
	mapval := NewMap()

	for i := 0; i < 10; i++ {
		go func() {
			val := mapval.Rd("key", time.Second*6)
			log.Println("读取值为->", val)
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			mapval.Out("key", val)
		}(i)
	}

	time.Sleep(time.Second * 10)
}

// 测试高并发读写，写多
func TestMultiRWMap(t *testing.T) {
	MockMultiVisit()
}
