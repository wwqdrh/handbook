package reflect

import "testing"

func TestIsStruct(t *testing.T) {
	type a struct {
		Name string
	}

	if !IsStruct(a{Name: "a"}) {
		t.Error()
	}
	if IsStruct(&a{Name: "a"}) {
		t.Error()
	}
	if GetStructName(a{Name: "a"}) != "a" {
		t.Error()
	}
}

func TestIterStructField(t *testing.T) {
	type User struct {
		ID   string
		Name string
	}
	if err := IterStructField(User{"1", "user"}); err != nil {
		t.Error(err)
	}
}
