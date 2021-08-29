package main

import "wwqdrh/handbook/cookbook/base/native/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
