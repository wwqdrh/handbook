package cookbook

import (
	"fmt"
	"unsafe"
)

// type hmap struct {
//     // 元素个数，调用 len(map) 时，直接返回此值
// 	count     int
// 	flags     uint8
// 	// buckets 的对数 log_2
// 	B         uint8
// 	// overflow 的 bucket 近似数
// 	noverflow uint16
// 	// 计算 key 的哈希的时候会传入哈希函数
// 	hash0     uint32
//     // 指向 buckets 数组，大小为 2^B
//     // 如果元素个数为0，就为 nil
// 	buckets    unsafe.Pointer
// 	// 等量扩容的时候，buckets 长度和 oldbuckets 相等
// 	// 双倍扩容的时候，buckets 长度会是 oldbuckets 的两倍
// 	oldbuckets unsafe.Pointer
// 	// 指示扩容进度，小于此地址的 buckets 迁移完成
// 	nevacuate  uintptr
// 	extra *mapextra // optional fields
// }
func reflectMap() {
	mp := make(map[string]int) // *hamp
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}

// type slice struct {
// 	array unsafe.Pointer // 指针，指向底层的数组，可以被同时引用，也就是说如果两个切片引用了同一地址空间的数据，是会相互影响的
// 	len int // 长度，打印时候只会打印len长度的元素，即使底层数组不止这么多
// 	cap int // 容量
// }
func reflectSlice() {
	s := make([]int, 9, 20) // slice
	Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))

	fmt.Println(Len, len(s)) // 9 9

	Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20
}

// zero copy
// string与slice的底层结构
// type StringHeader struct {
// 	Data uintptr
// 	Len  int
// }
// type SliceHeader struct {
// 	Data uintptr
// 	Len  int
// 	Cap  int
// }
func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
