package main

import "wwqdrh/handbook/cookbook/base/system/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}
