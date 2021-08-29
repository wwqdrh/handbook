package main

import (
	consul "github.com/hashicorp/consul/api"
	"wwqdrh/handbook/tools/micro0/discovery"
)

func main() {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"

	// faked name and port for example
	cli, err := discovery.NewClient(config, "localhost", "discovery", 8080)
	if err != nil {
		panic(err)
	}

	if err := discovery.Exec(cli); err != nil {
		panic(err)
	}
}
