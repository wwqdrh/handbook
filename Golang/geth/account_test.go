package geth

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddress(t *testing.T) {
	add := Address("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	if strings.ToLower(add.Hex()) != "0x71c7656ec7ab88b098defb751b7401b5f6d8976f" {
		t.Error("地址转换失败")
	}
}

func TestBalance(t *testing.T) {
	if val, err := Balance("0x71c7656ec7ab88b098defb751b7401b5f6d8976f"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(val)
	}
}

func TestWallet(t *testing.T) {
	if wall, err := Wallet(); err != nil {
		t.Error(err)
	} else {
		t.Log(wall)
	}
}

func TestKeyStore(t *testing.T) {
	// var p string = "./tmp"
	// if e := os.Mkdir(p, os.ModePerm); e != nil {
	// 	t.Error(e)
	// }

	// _, e := CreateKeyStore(p)
	// if e != nil {
	// 	t.Error(e)
	// }

	if _, e := ImportKeyStore("./tmp/UTC--2022-03-21T06-56-55.009904000Z--9ea3fe8f0a4ecd69c867c94047babfcc96b41e4e"); e != nil {
		t.Error(e)
	}
}
