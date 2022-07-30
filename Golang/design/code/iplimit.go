package code

import (
	"sync"
	"time"
)

// 在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。

// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
type RateLimit struct {
	LastVisit *sync.Map
}

func NewRateLimit() *RateLimit {
	return &RateLimit{
		LastVisit: &sync.Map{},
	}
}

func (l *RateLimit) Do(ip string) bool {
	val, loaded := l.LastVisit.LoadOrStore(ip, time.Now())
	if !loaded {
		return true
	}

	valTime, ok := val.(time.Time)
	if !ok {
		return false
	}
	if time.Since(valTime) > 3*time.Minute {
		l.LastVisit.Store(ip, time.Now())
		return true
	}
	return false
}
