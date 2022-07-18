package projecttool

import (
	"fmt"
	"testing"
)

var gs string

// BenchmarkSprint tests all the Sprint related benchmarks as sub benchmarks.
func BenchmarkSprintSub(b *testing.B) {
	b.Run("none", benchSprint)
	b.Run("format", benchSprintf)
}

// benchSprint tests the performance of using Sprint.
func benchSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// benchSprintf tests the performance of using Sprintf.
func benchSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
