package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/native/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
