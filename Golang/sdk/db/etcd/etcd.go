package etcd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdOptions struct {
	Endpoints   []string
	DialTimeout time.Duration
	UserName    string
	Password    string
}

type EtcdDriver struct {
	client *clientv3.Client
}

func NewEtcdDriver(options *EtcdOptions) (*EtcdDriver, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   options.Endpoints,   // []string{"localhost:2379"},
		DialTimeout: options.DialTimeout, // 5 * time.Second,
		Username:    options.UserName,    //   "root",
		Password:    options.Password,    //   "123456",
	})
	if err != nil {
		// handle error!
		return nil, err
	}
	return &EtcdDriver{
		client: cli,
	}, nil
}

func (d *EtcdDriver) Close() error {
	return d.client.Close()
}

func (d *EtcdDriver) HandleError(err error) error {
	if err != nil {
		switch err {
		case context.Canceled:
			return fmt.Errorf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			return fmt.Errorf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			return fmt.Errorf("client-side error: %v", err)
		default:
			return fmt.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	return nil
}

func (d *EtcdDriver) Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := d.client.Put(ctx, key, value)
	return d.HandleError(err)
}

func (d *EtcdDriver) PutWithLease(key, value string, ttl int64) (clientv3.LeaseID, error) {
	resp, err := d.client.Grant(context.TODO(), ttl)
	if err := d.HandleError(err); err != nil {
		return -1, err
	}

	_, err = d.client.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err := d.HandleError(err); err != nil {
		return -1, err
	}

	return resp.ID, nil
}

func (d *EtcdDriver) KeepAlive(id clientv3.LeaseID) (int64, error) {
	resp, err := d.client.KeepAliveOnce(context.TODO(), id)
	if err := d.HandleError(err); err != nil {
		return -1, err
	}
	return resp.TTL, nil
}

func (d *EtcdDriver) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := d.client.Get(ctx, key)
	if err := d.HandleError(err); err != nil {
		return "", err
	}
	for _, ev := range resp.Kvs {
		if string(ev.Key) == key {
			return string(ev.Value), nil
		}
	}
	return "", errors.New("未找到元素")
}

func (d *EtcdDriver) Watch(ctx context.Context, key string) error {
	rch := d.client.Watch(context.Background(), key, clientv3.WithPrefix())
	for {
		select {
		case wresp := <-rch:
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		case <-ctx.Done():
			fmt.Println("退出watch")
			return nil
		}
	}
}
