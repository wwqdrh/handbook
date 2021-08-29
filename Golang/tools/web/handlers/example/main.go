package main

import (
	"fmt"
	"net/http"

	"wwqdrh/handbook/tools/web/handlers"
)

func main() {
	http.HandleFunc("/name", handlers.HelloHandler)
	http.HandleFunc("/greeting", handlers.GreetingHandler)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
