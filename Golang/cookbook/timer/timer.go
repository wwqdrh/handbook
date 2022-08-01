package timer

import "time"

// 计时器

// time.AfterFunc
func fn2(msg interface{}, subscribers []chan interface{}) {
	count := 100
	concurrency := 1

	//采用Timer 而不是使用time.After 原因：time.After会产生内存泄漏 在计时器触发之前，垃圾回收器不会回收Timer
	pub := func(start int) {
		idleDuration := 5 * time.Millisecond
		idleTimeout := time.NewTimer(idleDuration)
		defer idleTimeout.Stop()
		for j := start; j < count; j += concurrency {
			if !idleTimeout.Stop() {
				// .Stop 如果已经停止或者过期返回false
				select {
				case <-idleTimeout.C:
				default:
				}
			}

			idleTimeout.Reset(idleDuration) // 如果已经触发了则返回true，如果过期或者停止返回false
			select {
			case subscribers[j] <- msg:
			case <-idleTimeout.C:
			}
		}
	}
	for i := 0; i < concurrency; i++ {
		go pub(i)
	}
}
