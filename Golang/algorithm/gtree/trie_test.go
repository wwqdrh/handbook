package gtree

import (
	"fmt"
	"testing"
)

// TODO error
func TestTrieSearch(t *testing.T) {
	filt := Constructor([]string{"apple"})
	fmt.Println(filt.F("a", "e"))

	filt = Constructor([]string{"abbba", "abba"})
	fmt.Println(filt.F("ab", "ba"))
}

// ok
func TestTrieSearch2(t *testing.T) {
	filt := Constructor2([]string{"apple"})
	fmt.Println(filt.F("a", "e"))

	filt = Constructor2([]string{"abbba", "abba"})
	fmt.Println(filt.F("ab", "ba"))
}
