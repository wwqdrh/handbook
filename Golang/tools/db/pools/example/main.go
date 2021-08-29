package main

import "wwqdrh/handbook/tools/db/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
