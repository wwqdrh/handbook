package main

import (
	"_examples/go-kit/util"
	"log"
)

func genRAS() {
	err := util.GenRSAPubAndPri(1024, "./pem")
	if err != nil {
		log.Fatal(err)
	}
}
