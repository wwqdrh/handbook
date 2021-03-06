package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

const (
	conRead  = 100
	conWrite = 100
	num      = 1000 // 写的个数
)

// testing.go:1152: race detected during execution of test
// --- FAIL: TestSafeSliceV1 (0.53s)
// === CONT
//     testing.go:1152: race detected during execution of test
// go test ./concurrent -v -run ^TestSafeSliceV1$ -race
func TestSafeSliceV1(t *testing.T) {
	safeSlice := NewSafeSlice()
	wait := sync.WaitGroup{}
	wait.Add(conRead + conWrite)

	for i := 0; i < conWrite; i++ {
		go func() {
			defer wait.Done()
			for cur := 1; cur <= num; cur++ {
				safeSlice.SafeSliceAppend(cur)
			}
		}()
	}
	for i := 0; i < conRead; i++ {
		go func() {
			defer wait.Done()
			for c := 0; c < 5; c++ {
				// fmt.Printf("第%d次执行", c)
				safeSlice.SafeSliceReadV1()
			}
		}()
	}
	wait.Wait()
	fmt.Println("执行结束")
}

// testing.go:1152: race detected during execution of test
// --- FAIL: TestSafeSliceV2 (0.53s)
// === CONT
//     testing.go:1152: race detected during execution of test
// go test ./concurrent -v -run ^TestSafeSliceV2$ -race
func TestSafeSliceV2(t *testing.T) {
	safeSlice := NewSafeSlice()
	wait := sync.WaitGroup{}
	wait.Add(conRead + conWrite)

	for i := 0; i < conWrite; i++ {
		go func() {
			defer wait.Done()
			for cur := 1; cur <= num; cur++ {
				safeSlice.SafeSliceAppend(cur)
			}
		}()
	}
	for i := 0; i < conRead; i++ {
		go func() {
			defer wait.Done()
			for c := 0; c < 5; c++ {
				// fmt.Printf("第%d次执行", c)
				safeSlice.SafeSliceReadV2()
			}
		}()
	}
	wait.Wait()
	fmt.Println("执行结束")
}

// go test ./concurrent -v -run ^TestSafeSliceV2_1$ -race
func TestSafeSliceV2_1(t *testing.T) {
	safeSlice := NewSafeSlice()
	wait := sync.WaitGroup{}
	wait.Add(conRead + conWrite)

	for i := 0; i < conWrite; i++ {
		go func() {
			defer wait.Done()
			for cur := 1; cur <= num; cur++ {
				safeSlice.SafeSliceAppend(cur)
			}
		}()
	}
	for i := 0; i < conRead; i++ {
		go func() {
			defer wait.Done()
			for c := 0; c < 5; c++ {
				// fmt.Printf("第%d次执行", c)
				safeSlice.SafeSliceReadV2_1()
			}
		}()
	}
	wait.Wait()
	fmt.Println("执行结束")
}

// testing.go:1152: race detected during execution of test
// --- FAIL: TestSafeSliceV3 (0.53s)
// === CONT
//     testing.go:1152: race detected during execution of test
// go test ./concurrent -v -run ^TestSafeSliceV3$ -race
func TestSafeSliceV3(t *testing.T) {
	safeSlice := NewSafeSlice()
	wait := sync.WaitGroup{}
	wait.Add(conRead + conWrite)

	for i := 0; i < conWrite; i++ {
		go func() {
			defer wait.Done()
			for cur := 1; cur <= num; cur++ {
				safeSlice.SafeSliceAppend(cur)
			}
		}()
	}
	for i := 0; i < conRead; i++ {
		go func() {
			defer wait.Done()
			for c := 0; c < 5; c++ {
				// fmt.Printf("第%d次执行", c)
				safeSlice.SafeSliceReadV3()
			}
		}()
	}
	wait.Wait()
	fmt.Println("执行结束")
}

// --- PASS: TestSafeSliceV4 (0.52s)
// PASS
// go test ./concurrent -v -run ^TestSafeSliceV4$ -race
func TestSafeSliceV4(t *testing.T) {
	safeSlice := NewSafeSliceByNode()
	wait := sync.WaitGroup{}
	wait.Add(conRead + conWrite)

	for i := 0; i < conWrite; i++ {
		go func() {
			defer wait.Done()
			for cur := 1; cur <= num; cur++ {
				safeSlice.SafeSliceAppend(cur)
			}
		}()
	}
	for i := 0; i < conRead; i++ {
		go func() {
			defer wait.Done()
			for c := 0; c < 5; c++ {
				// fmt.Printf("第%d次执行", c)
				safeSlice.SafeSliceRead()
			}
		}()
	}
	wait.Wait()
	fmt.Println("执行结束")
}

// channel接收并添加元素
func TestSafeSliceV5(t *testing.T) {
	safeSlice := NewSafeSliceByNode()
	wait := sync.WaitGroup{}
	wait.Add(conRead + conWrite)

	for i := 0; i < conWrite; i++ {
		go func() {
			defer wait.Done()
			for cur := 1; cur <= num; cur++ {
				safeSlice.SafeSliceAppendV2(cur)
			}
		}()
	}
	for i := 0; i < conRead; i++ {
		go func() {
			defer wait.Done()
			for c := 0; c < 5; c++ {
				// fmt.Printf("第%d次执行", c)
				safeSlice.SafeSliceRead()
			}
		}()
	}
	wait.Wait()
	safeSlice.Close()
	fmt.Println("执行结束")
}
