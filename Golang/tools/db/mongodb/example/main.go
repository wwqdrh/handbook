package main

import "wwqdrh/handbook/tools/db/mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
