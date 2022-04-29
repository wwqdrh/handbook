package bucket

// 滑动窗口算法, 与令牌桶一样，有应对突发流量的能力
// 滑动时间窗口算法，是从对普通时间窗口计数的优化。
// 使用普通时间窗口时，我们会为每个user_id/ip维护一个KV: uidOrIp: timestamp_requestCount。假设限制1秒1000个请求，那么第100ms有一个请求，这个KV变成 uidOrIp: timestamp_1，递200ms有1个请求，我们先比较距离记录的timestamp有没有超过1s，如果没有只更新count，此时KV变成 uidOrIp: timestamp_2。当第1100ms来一个请求时，更新记录中的timestamp并重置计数，KV变成 uidOrIp: newtimestamp_1
// 普通时间窗口有一个问题，假设有500个请求集中在前1s的后100ms，500个请求集中在后1s的前100ms，其实在这200ms没就已经请求超限了，但是由于时间窗每经过1s就会重置计数，就无法识别到此时的请求超限。
// 对于滑动时间窗口，我们可以把1ms的时间窗口划分成10个time slot, 每个time slot统计某个100ms的请求数量。每经过100ms，有一个新的time slot加入窗口，早于当前时间100ms的time slot出窗口。窗口内最多维护10个time slot，储存空间的消耗同样是比较低的。
// 主要就是实现sliding window算法。可以参考Bilibili开源的kratos框架里circuit breaker用循环列表保存time slot对象的实现，
// 他们这个实现的好处是不用频繁的创建和销毁time slot对象。下面给出一个简单的基本实现：
import (
	"sync"
	"time"
)

var winMu map[string]*sync.RWMutex

func init() {
	winMu = make(map[string]*sync.RWMutex)
}

type timeSlot struct {
	timestamp time.Time // 这个timeSlot的时间起点
	count     int       // 落在这个timeSlot内的请求数
}

func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

type SlidingWindowLimiter struct {
	SlotDuration time.Duration // time slot的长度
	WinDuration  time.Duration // sliding window的长度
	numSlots     int           // window内最多有多少个slot
	windows      map[string][]*timeSlot
	maxReq       int // win duration内允许的最大请求数
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration:  winDuration,
		numSlots:     int(winDuration / slotDuration),
		windows:      make(map[string][]*timeSlot),
		maxReq:       maxReq,
	}
}

// 获取user_id/ip的时间窗口
func (l *SlidingWindowLimiter) getWindow(uidOrIp string) []*timeSlot {
	win, ok := l.windows[uidOrIp]
	if !ok {
		win = make([]*timeSlot, 0, l.numSlots)
	}
	return win
}

func (l *SlidingWindowLimiter) storeWindow(uidOrIp string, win []*timeSlot) {
	l.windows[uidOrIp] = win
}

func (l *SlidingWindowLimiter) validate(uidOrIp string) bool {
	// 同一user_id/ip并发安全
	mu, ok := winMu[uidOrIp]
	if !ok {
		var m sync.RWMutex
		mu = &m
		winMu[uidOrIp] = mu
	}
	mu.Lock()
	defer mu.Unlock()

	win := l.getWindow(uidOrIp)
	now := time.Now()
	// 已经过期的time slot移出时间窗
	timeoutOffset := -1
	for i, ts := range win {
		if ts.timestamp.Add(l.WinDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		win = win[timeoutOffset+1:]
	}

	// 判断请求是否超限
	var result bool
	if countReq(win) < l.maxReq {
		result = true
	}

	// 记录这次的请求数
	var lastSlot *timeSlot
	if len(win) > 0 {
		lastSlot = win[len(win)-1]
		if lastSlot.timestamp.Add(l.SlotDuration).Before(now) {
			lastSlot = &timeSlot{timestamp: now, count: 1}
			win = append(win, lastSlot)
		} else {
			lastSlot.count++
		}
	} else {
		lastSlot = &timeSlot{timestamp: now, count: 1}
		win = append(win, lastSlot)
	}

	l.storeWindow(uidOrIp, win)

	return result
}

func (l *SlidingWindowLimiter) getUidOrIp() string {
	return "127.0.0.1"
}

func (l *SlidingWindowLimiter) IsLimited() bool {
	return !l.validate(l.getUidOrIp())
}
