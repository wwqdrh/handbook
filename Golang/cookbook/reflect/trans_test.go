package reflect

import (
	"fmt"
	"testing"
)

func TestBytes2String(t *testing.T) {
	a := bytes2string([]byte("hello"))
	fmt.Println(a)
	fmt.Println(len(a))
}
