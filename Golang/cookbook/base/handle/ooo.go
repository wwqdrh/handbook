package handle

type rect struct {
	width, height int
}

// 这里的 `area` 是一个拥有 `*rect` 类型接收器(receiver)的方法。
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收者定义方法。
// 这是一个值类型接收者的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
