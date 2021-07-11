package utils

func IntSliceCompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
