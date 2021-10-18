package algorithm

import (
	"fmt"
	"math"
)

////////////////////
// 公共函数
////////////////////

func getMaxInArr(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

//交换
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

////////////////////
// 非比较类排序算法
////////////////////

//计数排序: 找到最大值 按照值来创建桶
func CountSort(nums []int) (res []int) {
	maxVal := getMaxInArr(nums)
	bucket := make([]int, maxVal+1)
	for _, num := range nums {
		bucket[num]++
	}

	for i, cnt := range bucket {
		for ; cnt > 0; cnt-- {
			res = append(res, i)
		}
	}
	return
}

// 桶排序 按照元素个数建桶 在求元素的所在的桶编号的时候 需要 num * (n - 1) / maxNum 表示一个值占多少个桶乘上num就能知道在哪个桶
// 按照桶的粒度是有序的 然后再在桶内进行排序
func BucketSort(nums []int) (res []int) {
	buckets := make([][]int, len(nums))

	maxVal, length := getMaxInArr(nums), len(nums)
	for _, num := range nums {
		index := num * (length - 1) / maxVal
		buckets[index] = append(buckets[index], num)
	}

	// 对桶内进行排序
	for _, bucket := range buckets {
		cur := InsertionSort(bucket)
		res = append(res, cur...)
	}
	return
}

// 基数排序
// 最高位优先(Most Significant Digit first)法，简称MSD法：
// 先按k1排序分组，同一组中记录，关键码k1相等，再对各组按k2排序分成子组，之后，对后面的关键码继续这样的排序分组，直到按最次位关键码kd对各子组排序后。再将各组连接起来，便得到一个有序序列。
// 最低位优先(Least Significant Digit first)法，简称LSD法：
// 先从kd开始排序，再对kd-1进行排序，依次重复，直到对k1排序后便得到一个有序序列。
func RadixSort(nums []int) []int {
	a := make([]int, len(nums))
	copy(a, nums)

	max := getMaxInArr(a)
	//获取最大值的位数
	var count int = len(fmt.Sprint(max))

	//给桶中对应的位置放数据
	for i := 0; i < count; i++ {

		theData := int(math.Pow10(i)) //10的i次方
		//建立并初始化空桶
		var bucket [10][10]int
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				bucket[i][j] = -1
			}
		}
		//给桶赋值
		for k := 0; k < len(a); k++ {
			theResidue := (a[k] / theData) % 10 //取余
			for m := 0; m < 10; m++ {
				if bucket[theResidue][m] == -1 {
					bucket[theResidue][m] = a[k]
					break
				} else {
					continue
				}
			}
		}
		var x = 0
		//出桶
		for p := 0; p < len(bucket); p++ {
			for q := 0; q < len(bucket[p]); q++ {
				if bucket[p][q] != -1 {
					a[x] = bucket[p][q]
					x++
				} else {
					break
				}
			}
		}
	}
	return a
}

// 冒泡排序
// 每一趟排序把最大的移到末尾
func BubbleSort(nums []int, cb func(i, j int) bool) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		f := true // 如果一趟下来是有序的
		for j := 0; j < length-i-1; j++ {
			if !cb(j, j+1) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				f = false
			}
		}
		if f {
			break
		}
	}
	return nums
}

// 快速排序
// 分治的思想 将比基数大的放在右边 比基数小的放在左边 然后递归处理
func QuickSort(nums []int, low, high int) []int {
	if low >= high {
		return nums
	}

	start := nums[low]
	i := low
	for j := low + 1; j <= high; j++ {
		if nums[j] <= start {
			i++
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	nums[i], nums[low] = nums[low], nums[i]
	QuickSort(nums, low, i-1)
	QuickSort(nums, i+1, high)
	return nums
}

// 插入排序
// 每一步将需要排序的元素插入到前面已经有序的数列中
func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > cur; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = cur
	}
	return nums
}

// 希尔排序 简单插入排序的改进 将待排序的数组按照一定增量进行分组
// 跳跃式分组 对每组中的元素进行排序
// 随着增量逐渐变为1那么就是对整个数组进行排序 这时数组已经有序
func ShellSort(nums []int) []int {
	length := len(nums)
	for gap := len(nums) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			for j := i; j-gap >= 0 && nums[j] < nums[j-gap]; j -= gap {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
			}
		}
	}
	return nums
}

// 选择排序
// 第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，
// 然后再从剩余的未排序元素中寻找到最小（大）元素
func SelectionSort(a []int) []int {
	l := len(a)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

//堆排序
func heapSort(a []int) []int {
	arrLen := len(a)
	buildMaxHeap(a, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(a, 0, i)
		arrLen -= 1
		heapify(a, 0, arrLen)
	}
	return a
}

//建立大根堆
func buildMaxHeap(a []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(a, i, arrLen)
	}
}

func heapify(a []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && a[left] > a[largest] {
		largest = left
	}
	if right < arrLen && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		swap(a, i, largest)
		heapify(a, largest, arrLen)
	}
}

//归并排序
func MergeSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	middle := length / 2
	left := a[0:middle]
	right := a[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

//归并
func merge(left []int, right []int) []int {
	var result []int
	leftIdx, rightIdx := 0, 0
	for (leftIdx < len(left)) && (rightIdx < len(right)) {
		if left[leftIdx] <= right[rightIdx] {
			result = append(result, left[leftIdx])
			leftIdx++
		} else {
			result = append(result, right[rightIdx])
			rightIdx++
		}
	}
	if leftIdx < len(left) {
		result = append(result, left[leftIdx:]...)
	}
	if rightIdx < len(right) {
		result = append(result, right[rightIdx:]...)
	}
	return result
}
