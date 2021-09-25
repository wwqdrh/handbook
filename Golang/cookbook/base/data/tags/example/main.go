package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/data/tags"
)

func main() {

	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
