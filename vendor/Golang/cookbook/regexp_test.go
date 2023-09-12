package cookbook

import "testing"

func TestMostCommonWord(t *testing.T) {
	mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
}
