package basic

import (
	"log"
	"net"
	"net/rpc"
)

type MathService struct {
}

type Args struct {
	A, B int
}

func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func ServerStart() {
	rpc.RegisterName("MathService", new(MathService))
	l, err := net.Listen("tcp", ":8088") //注意 “：” 不要忘了写
	if err != nil {
		log.Fatal("listen error", err)
	}
	rpc.Accept(l)
}
