package sequence

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	l := NewSkiplist()
	l.Add(100)
	l.Add(50)
	l.Add(99)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.val)
	}

	fmt.Println(l.Search(99))
	fmt.Println(l.Search(98))
}
