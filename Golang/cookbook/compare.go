package cookbook

import "fmt"

func comparePointer() {
	var a *int
	var b *string
	fmt.Println(a == nil, b == nil)
	// fmt.Println(a == b)  // type和值都必须是nil
}
