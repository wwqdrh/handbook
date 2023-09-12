package main

import (
	"net/http"

	"wwqdrh/handbook/tools/micro0/consensus"
)

func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
