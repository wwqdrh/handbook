package cookbook

import (
	"fmt"
	"testing"
)

func TestSimpleRand(t *testing.T) {
	SimpleRand()
	fmt.Println("===with seed===")
	SeedRand()
}
