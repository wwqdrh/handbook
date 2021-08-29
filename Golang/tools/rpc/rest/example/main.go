package main

import "wwqdrh/handbook/tools/rpc/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
