package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/native/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
