package mutex

import (
	"fmt"
	"net/url"
	"sync"
	"testing"
)

// vet error, 不能赋值锁
func TestCopy1(t *testing.T) {
	var amux sync.Mutex
	b := amux
	b.Lock()
	b.Unlock()
}

// 会发生复制的情况
// struct 类型变量的值传递作为返回值
// struct 类型变量的值传递作为 receiver
func TestCopy2(t *testing.T) {
	type URL struct {
		Ip     string
		Port   string
		mux    sync.RWMutex
		params url.Values
	}

	var url1 URL
	url2 := url1
	fmt.Println(url2)
}

var age int
var wg sync.WaitGroup

type Person struct {
	mux sync.Mutex
}

func (p Person) AddAge() {
	defer wg.Done()
	p.mux.Lock()
	age++
	defer p.mux.Unlock()

}
func TestCopy3(t *testing.T) {
	for i := 0; i < 100; i++ {
		p1 := Person{
			mux: sync.Mutex{},
		}
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go p1.AddAge()
		}
		wg.Wait()
		fmt.Println(age)
		age = 0
	}
}

type Person2 struct {
	mux sync.Mutex
}

func Reduce(p Person2) {
	fmt.Println("step...")
	p.mux.Lock()
	fmt.Println(p)
	defer p.mux.Unlock()
	fmt.Println("over...")
}

func TestCopy4(t *testing.T) {
	var p Person2
	p.mux.Lock()
	go Reduce(p) // 加加了锁的复制进去(状态也复制进去了，再次加锁会死锁)
	p.mux.Unlock()
	fmt.Println(111)
	for {
	}
}

func TestDeadLockCantFound(t *testing.T) {
	l := sync.Mutex{}

	l.Lock()
	l.Lock()
}

func TestRMMutex(t *testing.T) {
	l := sync.RWMutex{}
	l.Lock()
	defer l.Unlock()

	a := func() {
		l.RLock()
		defer l.Unlock()
	}

	a()
}
