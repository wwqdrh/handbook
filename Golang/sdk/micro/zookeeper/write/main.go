package main

import (
	"fmt"
	"time"
	"wwqdrh/handbook/scripts/middleware/zookeeper"
)

func main() {
	zkManager := zookeeper.NewZkManager([]string{"127.0.0.1:2181"})
	zkManager.GetConnect()
	defer zkManager.Close()
	i := 0

	for {
		conf := fmt.Sprintf("{name:" + fmt.Sprint(i) + "}")
		zkManager.SetPathData("/rs_server_conf", []byte(conf), int32(i))
		time.Sleep(5 * time.Second)
		i++
	}
}
