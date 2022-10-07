package cookbook

import "testing"

func TestSliceDo(t *testing.T) {
	DoSlice(4)
}

func TestSliceDo2(t *testing.T) {
	DoSlice2()
}

func TestDoSlice(t *testing.T) {
	DoSlice3(3)
}

func TestArr2Slice(t *testing.T) {
	arr2slice()
}

func TestSliceSort(t *testing.T) {
	slicesort()
}
