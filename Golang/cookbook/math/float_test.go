package math

import "math"

func ExampleIsInt() {
	IsInt(math.Float32bits(3.0), 127)

	// Output: INTEGER
}
