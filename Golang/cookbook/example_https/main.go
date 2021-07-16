package main

import (
	"flag"
	"fmt"
	"log"

	"example_https/server"
)

func main() {
	type_ := flag.String("type", "", "类型")

	flag.Parse()
	if *type_ == "" {
		log.Fatal("未指定类型")
	} else if *type_ == "server" {
		server.Server()
	} else if *type_ == "client" {
		res, _ := server.BcjClient([]string{"hello", "world", "hello"})
		fmt.Printf("%#v", res)
	}
}
