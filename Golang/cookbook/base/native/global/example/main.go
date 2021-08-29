package main

import "wwqdrh/handbook/cookbook/base/native/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
