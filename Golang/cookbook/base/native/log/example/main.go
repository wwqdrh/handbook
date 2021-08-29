package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/native/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
