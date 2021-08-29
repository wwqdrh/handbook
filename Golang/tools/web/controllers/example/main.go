package main

import (
	"fmt"
	"net/http"

	"wwqdrh/handbook/tools/web/controllers"
)

func main() {
	storage := controllers.MemStorage{}
	c := controllers.New(&storage)
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
