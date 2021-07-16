package server

import (
	"log"
	"os"
	"testing"
)

// 启动服务，并且进行多次客户端测试来验证正确性
func TestHttpsServerClient(t *testing.T) {
	os.Chdir("../")

	// 启动服务端
	go func() {
		Server()
	}()

	wait := make(chan bool)
	// 启动客户端，并且进行多次测试
	go func(wait chan bool) {
		defer func() { wait <- true }()

		// 第一次测试
		var res, expect []bool
		var err error

		res, err = BcjClient([]string{"hello", "world", "hello"})
		if err != nil {
			log.Println(err)
			// return
		}
		expect = []bool{false, false, true}
		for i := 0; i < len(expect); i++ {
			if res[i] != expect[i] {
				t.Error("类型不对")
			}
		}

		// 第二次测试
		res, err = BcjClient([]string{"hello", "world", "hello"})
		if err != nil {
			log.Println(err)
			// return
		}
		expect = []bool{true, true, true}
		for i := 0; i < len(expect); i++ {
			if res[i] != expect[i] {
				t.Error("类型不对")
			}
		}
	}(wait)

	<-wait
}

func TestHttpsServerClientWithClear(t *testing.T) {
	os.Chdir("../")

	// 启动服务端
	go func() {
		Server()
	}()

	wait := make(chan bool)
	// 启动客户端，并且进行多次测试
	go func(wait chan bool) {
		defer func() { wait <- true }()

		var res, expect []bool
		var err error

		// 第一次测试
		err = ClearCache()
		if err != nil {
			t.Error("清理缓存失败")
			return
		}

		res, err = BcjClient([]string{"hello", "world", "hello"})
		if err != nil {
			log.Println(err)
			// return
		}
		expect = []bool{false, false, true}
		for i := 0; i < len(expect); i++ {
			if res[i] != expect[i] {
				t.Error("类型不对")
			}
		}

		// 第二次测试
		err = ClearCache()
		if err != nil {
			t.Error("清理缓存失败")
			return
		}

		res, err = BcjClient([]string{"hello", "world", "hello"})
		if err != nil {
			log.Println(err)
			// return
		}
		expect = []bool{false, false, true}
		for i := 0; i < len(expect); i++ {
			if res[i] != expect[i] {
				t.Error("类型不对")
			}
		}
	}(wait)

	<-wait
}
