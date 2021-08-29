package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/native/tags"
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
