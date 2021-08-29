package main

import "wwqdrh/handbook/tools/db/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
