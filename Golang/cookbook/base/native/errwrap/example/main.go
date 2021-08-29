package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/native/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
