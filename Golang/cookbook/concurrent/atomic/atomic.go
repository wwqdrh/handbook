package atomic

import "sync/atomic"

var res int64

func BasicWrite() {
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&res, 1)
	}
}

func BasicRead() int64 {
	return atomic.LoadInt64(&res)
}
