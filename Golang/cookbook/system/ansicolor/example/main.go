package main

import (
	"fmt"

	"wwqdrh/handbook/cookbook/base/system/ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red!",
	}

	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "Now I'm green!"

	fmt.Println(r.String())

	r.TextColor = ansicolor.ColorNone
	r.Text = "Back to normal..."

	fmt.Println(r.String())
}
