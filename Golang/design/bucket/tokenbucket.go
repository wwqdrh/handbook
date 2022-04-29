package bucket

import (
	"sync"
	"time"
)

// 想象有一个木桶，以固定的速度往木桶里加入令牌，木桶满了则不再加入令牌。服务收到请求时尝试从木桶中取出一个令牌，如果能够得到令牌则继续执行后续的业务逻辑；如果没有得到令牌，直接返回反问频率超限的错误码或页面等，不继续执行后续的业务逻辑
// 适合电商抢购或者微博出现热点事件这种场景，因为在限流的同时可以应对一定的突发流量。如果采用均匀速度处理请求的算法，在发生热点时间的时候，会造成大量的用户无法访问，对用户体验的损害比较大。

// 假设每100ms生产一个令牌，按user_id/IP记录访问最近一次访问的时间戳 t_last 和令牌数，
// 每次请求时如果 now - last > 100ms, 增加 (now - last) / 100ms个令牌。然后，如果令牌数 > 0，令牌数 -1 继续执行后续的业务逻辑，否则返回请求频率超限的错误码或页面。

var (
	recordMu = make(map[string]*sync.RWMutex)
	mux      = sync.RWMutex{}
)

type (
	TokenBucket struct {
		BucketSize int // 木桶内的容量; 最多可以存放多少个令牌
		TokenRate  time.Duration
		records    map[string]*record
	}

	record struct {
		last  time.Time
		token int
	}
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewTokenBucket(bucketSize int, tokenRate time.Duration) *TokenBucket {
	return &TokenBucket{
		BucketSize: bucketSize,
		TokenRate:  tokenRate,
		records:    make(map[string]*record),
	}
}

// 获取请求用户的user_id或者ip地址
func (t *TokenBucket) getUidOrdIp() string {
	return "127.0.0.1"
}

// 保存user_id/ip最近一次请求的时间戳和令牌数量
func (t *TokenBucket) storeRecord(uidOrIp string, r *record) {
	t.records[uidOrIp] = r
}

func (t *TokenBucket) getRecord(uidOrIp string) *record {
	if r, ok := t.records[uidOrIp]; ok {
		return r
	}
	return &record{}
}

func (t *TokenBucket) Put(uidOrIp string) bool {
	// 并发修改同一个用户的记录上写锁
	// 双检测
	mux.RLock()
	rl, ok := recordMu[uidOrIp]
	mux.RUnlock()
	if !ok {
		mux.Lock()
		if _, ok := recordMu[uidOrIp]; !ok {
			rl = &sync.RWMutex{}
			recordMu[uidOrIp] = rl
		}
		mux.Unlock()
	}

	rl.Lock()
	defer rl.Unlock()
	r := t.getRecord(uidOrIp)
	now := time.Now()
	if r.last.IsZero() {
		// 第一次访问初始化最大令牌数
		r.last, r.token = now, t.BucketSize
	} else {
		if r.last.Add(t.TokenRate).Before(now) {
			// 如果与上次请求的间隔超过了token rate
			// 则增加令牌，更新last
			r.token += max(int(now.Sub(r.last)/t.TokenRate), t.BucketSize)
			r.last = now
		}
	}
	var result bool
	if r.token > 0 {
		// 如果令牌数大于1，取走一个令牌，put成功
		r.token--
		result = true
	}
	// 保存最新的record
	t.storeRecord(uidOrIp, r)
	return result
}

// 返回是否被限流
func (t *TokenBucket) IsLimited() bool {
	return !t.Put(t.getUidOrdIp())
}
