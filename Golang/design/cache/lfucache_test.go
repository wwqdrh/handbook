package cache

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	c := NewLFUCache(2)
	c.Put(2, 1)
	c.Put(1, 1)
	fmt.Println(c.Get(2))
	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
