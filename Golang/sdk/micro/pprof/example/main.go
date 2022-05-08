package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"wwqdrh/handbook/tools/micro3/pprof/crypto"
)

func main() {

	http.HandleFunc("/guess", crypto.GuessHandler)
	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}
