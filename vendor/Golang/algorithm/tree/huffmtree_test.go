package tree

import (
	"fmt"
	"testing"
)

func TestHuffmTree(t *testing.T) {
	keyList := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	// keyList2 := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	valueList := []float64{0.12, 0.4, 0.29, 0.90, 0.1, 1.1, 1.23, 0.01}

	hfmt, err := NewHuffmanTree(keyList, valueList)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hfmt.WPLValue(0))
}
