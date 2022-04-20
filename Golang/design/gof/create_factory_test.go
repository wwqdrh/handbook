package gof

import (
	"fmt"
	"testing"
)

func TestDI(t *testing.T) {

	// A 依赖关系 A -> B -> C
	// C C
	type C struct {
		Num int
	}
	// B B
	type B struct {
		C *C
	}
	type A struct {
		B *B
	}
	// NewA NewA
	NewA := func(b *B) *A {
		return &A{
			B: b,
		}
	}
	// NewB NewB
	NewB := func(c *C) *B {
		return &B{C: c}
	}
	// NewC NewC
	NewC := func() *C {
		return &C{
			Num: 1,
		}
	}

	container := NewDIContiner()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}
}
