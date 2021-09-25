package data

import (
	"crypto/sha1"
	"fmt"
)

func ShaExample() {
	s := "sha1 this string"

	// 产生一个散列值的方式是 `sha1.New()`，`sha1.Write(bytes)`，
	// 然后 `sha1.Sum([]byte{})`。这里我们从一个新的散列开始。
	h := sha1.New()

	// 写入要处理的字节。如果是一个字符串，
	// 需要使用 `[]byte(s)` 将其强制转换成字节数组。
	h.Write([]byte(s))

	// `Sum` 得到最终的散列值的字符切片。`Sum` 接收一个参数，
	// 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做。
	bs := h.Sum(nil)

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。
	// 我们这里也使用 `%x` 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
