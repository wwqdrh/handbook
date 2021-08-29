package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/concurrent/waitgroup"
)

func main() {
	sites := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://www.google.com/search?q=golang",
	}

	resps, err := waitgroup.Crawl(sites)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resps received:", resps)
}
