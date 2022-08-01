package basic

import (
	"fmt"
	"log"
	"net/rpc"
)

func ClientStart() {
	client, err := rpc.Dial("tcp", "localhost:8088")
	if err != nil {
		log.Fatal("dialing")
	}
	args := Args{A: 1, B: 2}
	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error", err)
	}
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
}
