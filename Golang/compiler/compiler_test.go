package compiler

import (
	"runtime"
	"testing"
	"time"
)

func TestNumFinizaler(t *testing.T) {
	numFinizaler()
	runtime.GC()
	print("第一次gc done \n")
	time.Sleep(10 * time.Second)
	runtime.GC() // 之前的清理操作还未执行完
	print("第二次gc done \n")
	time.Sleep(3 * time.Second)
}
