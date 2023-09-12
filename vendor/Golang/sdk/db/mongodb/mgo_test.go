package mongodb

import "testing"

func TestMgo(t *testing.T) {
	if err := Exec(); err != nil {
		panic(err)
	}
}
