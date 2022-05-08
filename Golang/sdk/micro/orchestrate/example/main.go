package main

import (
	mgo "gopkg.in/mgo.v2"
	"wwqdrh/handbook/tools/micro0/orchestrate"
)

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err)
	}
	if err := orchestrate.ConnectAndQuery(session); err != nil {
		panic(err)
	}
}
