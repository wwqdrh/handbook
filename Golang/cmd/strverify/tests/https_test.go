package tests

import (
	"context"
	"reflect"
	"strverify"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestStringVerify(t *testing.T) {
	url := "https://localhost:8080/stringverify"
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(1 * time.Second)
	}()

	strverify.HttpClientExample(ctx)

	var (
		res []bool
		err error
	)

	res, err = strverify.StringVerifyRequest(url, []string{"asda", "asda", "asdaso", "asdjakld"})
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(res, []bool{false, true, false, false}) {
		t.Error("TestStringVerify失败")
	}

	res, err = strverify.StringVerifyRequest(url, []string{"asda", "asda", "asdaso", "asdjakld"})
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(res, []bool{true, true, true, true}) {
		t.Error("TestStringVerify失败")
	}

	// 测试并发的正确性 多个相同字符串同时执行 但是只能有一个返回false其他的必须都是true
	var count int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(7)
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	go func() {
		res, err := strverify.StringVerifyRequest(url, []string{"asdfghjkl"})
		if err != nil {
			t.Error(err)
		} else if !res[0] {
			atomic.AddInt64(&count, 1)
		}
		wg.Done()
	}()
	wg.Wait()
	if count != 1 {
		t.Errorf("TestStringVerify失败 %d", count)
	}
}
