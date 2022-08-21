package cookbook

import (
	"fmt"
	"testing"
)

func TestComparePinter(t *testing.T) {
	comparePointer()
}

func TestCompareMap(t *testing.T) {
	fmt.Println(compareMap())
}

func TestCompareNil(t *testing.T) {
	compareNil()
}
