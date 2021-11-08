package design

import (
	"sync"
	"sync/atomic"
)

// 懒汉模式
// 1、双重检验: 需要原子操作 不然在多cpu情况下 还是可能会产生静态
// 2、sync.Once

// 双重验证
type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.m.Lock()
		defer o.m.Unlock()
		if o.done == 0 {
			defer atomic.StoreUint32(&o.done, 1)
			f()
		}
	}
}

// sync.Once模式
// 可以看到sync.Once内部其实也是一个双重检验锁，但是对于共享变量（done字段）的读和写使用了atomic包的StoreUint32和LoadUint32方法
// sync.Once使用一个32位无符号整数表示共享变量，即使是32位变量的读写操作都需要atomic包方法来实现原子性，更说明了go里边指针的读写不能保证原子性
var (
	instance2 *int
	once      sync.Once
)

func OnceInstance() *int {
	once.Do(func() {
		if instance2 == nil {
			i := 1
			instance2 = &i
		}
	})
	return instance2
}
