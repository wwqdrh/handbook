package geth

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	if c, err := Connection("https://cloudflare-eth.com"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(c)
	}
}
