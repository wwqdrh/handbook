package main

import "wwqdrh/handbook/tools/rpc/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
