package main

import (
	"fmt"
	"net"
	"wwqdrh/handbook/cookbook/network"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	network.Encode(conn, "hello world 0!!!")
}
