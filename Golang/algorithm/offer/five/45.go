package five

import (
	"sort"
	"strconv"
	"strings"
)

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

input: [10, 2]
output: "102"

input: [3, 30, 34, 5, 9]
output: "3033459"
*/

func Hand45(nums []int) string {
	numStr := make([]string, len(nums))
	for i, num := range nums {
		numStr[i] = strconv.FormatInt(int64(num), 10)
	}
	sort.Slice(numStr, func(i, j int) bool { return numStr[i]+numStr[j] < numStr[j]+numStr[i] })
	return strings.Join(numStr, "")
}
