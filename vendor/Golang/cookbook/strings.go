package cookbook

import "strings"

func getSentenceWord1(word string) []string {
	return strings.Fields(word)
}

func getSentenceWord2(word string) []string {
	return strings.Split(word, " ")
}
