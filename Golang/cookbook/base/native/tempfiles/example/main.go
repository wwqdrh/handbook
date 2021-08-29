package main

import "wwqdrh/handbook/tools/micro/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
