package cookbook

import (
	"fmt"
	"reflect"
)

func compareNil() {
	var a map[int]int
	fmt.Println(a == nil)

	var b interface{} = a
	fmt.Println(b == nil)
}

func comparePointer() {
	var a *int
	var b *string
	fmt.Println(a == nil, b == nil)
	// fmt.Println(a == b)  // type和值都必须是nil

	float1, float2 := 3.5, 3.5
	float3 := 7.0 / 2.0
	float4, float5 := 1.5, 1.3
	float6, float7 := 0.4, 0.2
	fmt.Println(float1 == float2)
	fmt.Println(float1 == float3)
	fmt.Println(float4-float5, float4-float5 == 0.2)
	fmt.Println(float6-float7, float6-float7 == 0.2)
	fmt.Println(float4+float5 == 2.8)

	var (
		nums1 float32
		nums2 float64
		nums3 float32
		nums4 float64
		nums5 float64
		nums6 float64
	)
	for i := 0; i < 7; i++ {
		nums1 += 0.1
		nums2 += 0.1
		nums3 += 0.2
		nums4 += 0.3
		nums5 += 0.4
		nums6 += 0.5
	}
	fmt.Println(nums1, nums2, nums3, nums4, nums5, nums6)

	var f1 float32 = 9.90
	fmt.Println(f1 * 100)

	var f2 float64 = 9.90
	fmt.Println(f2 * 100)
}

func compareMap() bool {
	a := map[string]interface{}{
		"a": 1,
		"b": 2,
	}

	b := map[string]interface{}{
		"b": 2,
		"a": 1,
	}

	return reflect.DeepEqual(a, b)
}
