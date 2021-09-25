package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/native/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}
