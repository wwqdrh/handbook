package etcd

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestBasicWatch(t *testing.T) {
	// cli, err := clientv3.New(clientv3.Config{
	// 	Endpoints:   []string{"localhost:2379"},
	// 	DialTimeout: 5 * time.Second,
	// 	Username:    "root",
	// 	Password:    "123456",
	// })
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer cli.Close()

	driver, err := NewEtcdDriver(&EtcdOptions{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		UserName:    "root",
		Password:    "123456",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Close()
	cli := driver.client

	// watches within ['foo1', 'foo4'), in lexicographical order
	// rch := cli.Watch(context.Background(), "foo1", clientv3.WithRange("foo4"))
	rch := cli.Watch(context.Background(), "/test", clientv3.WithPrefix())

	go func() {
		// cli.Put(context.Background(), "/test/crud/key1", "value1")
		driver.Put("/test/crud/key1", "value1")
		// cli.Put(context.Background(), "/test/crud/key2", "value2")
		driver.Put("/test/crud/key2", "value2")
		// cli.Put(context.Background(), "/test/crud/key3", "value3")
		driver.Put("/test/crud/key3", "value3")
		// cli.Put(context.Background(), "/test/crud/key4", "value4")
		driver.Put("/test/crud/key4", "value4")
	}()

	i := 0
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			i++
			if i == 4 {
				// After 3 messages we are done.
				cli.Delete(context.Background(), "foo", clientv3.WithPrefix())
				cli.Close()
				return
			}
		}
	}
}

// 租约撤销: revoke
// 租约续租: KeepAlive
func TestBasicLease(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "123456",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	// minimum lease TTL is 5-second
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		t.Error(err)
	}

	// after 5 seconds, the key 'foo' will be removed
	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		t.Error(err)
	}

	resp, err = cli.Grant(context.TODO(), 5)
	if err != nil {
		t.Error(err)
	}

	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		t.Error(err)
	}

	// to renew the lease only once
	ka, kaerr := cli.KeepAliveOnce(context.TODO(), resp.ID)
	if kaerr != nil {
		t.Error(kaerr)
	}

	fmt.Println("ttl:", ka.TTL)
}

func TestDriverGetPut(t *testing.T) {
	options := &EtcdOptions{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		UserName:    "root",
		Password:    "123456",
	}
	driver, err := NewEtcdDriver(options)
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Close()

	// 测试put
	if err := driver.Put("/test/crud/key1", "value1"); err != nil {
		t.Fatal(err)
	}

	// 测试get
	if res, err := driver.Get("/test/crud/key1"); err != nil || res != "value1" {
		t.Fatal(errors.New("查询失败"))
	}
}

func TestDriverWatchPut(t *testing.T) {
	options := &EtcdOptions{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		UserName:    "root",
		Password:    "123456",
	}
	driver, err := NewEtcdDriver(options)
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 协程执行watch
	go driver.Watch(ctx, "/test/crud")
	time.Sleep(1 * time.Second) //!!!必须保证前面先进行watch
	// rch := driver.client.Watch(context.Background(), "/test/crud", clientv3.WithPrefix())
	// !!! 必须将rch先执行才能监听到变化，如果放在后面就没有了！！！！

	// cli.Put(context.Background(), "/test/crud/key1", "value1")
	driver.Put("/test/crud/key1", "value1")
	// cli.Put(context.Background(), "/test/crud/key2", "value2")
	driver.Put("/test/crud/key2", "value2")
	// cli.Put(context.Background(), "/test/crud/key3", "value3")
	driver.Put("/test/crud/key3", "value3")
	// cli.Put(context.Background(), "/test/crud/key4", "value4")
	driver.Put("/test/crud/key4", "value4")
	time.Sleep(1 * time.Second) // 保证前面监听

	// i := 0
	// for wresp := range rch {
	// 	for _, ev := range wresp.Events {
	// 		fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	// 		i++
	// 		if i == 4 {
	// 			// After 3 messages we are done.
	// 			return
	// 		}
	// 	}
	// }

}
