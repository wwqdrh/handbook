package handle

import (
	"fmt"
	"os"
)

func DeferExample() {
	createFile := func(p string) *os.File {
		fmt.Println("creating")
		f, err := os.Create(p)
		if err != nil {
			panic(err)
		}
		return f
	}

	writeFile := func(f *os.File) {
		fmt.Println("writing")
		fmt.Fprintln(f, "data")

	}

	closeFile := func(f *os.File) {
		fmt.Println("closing")
		err := f.Close()
		// 关闭文件时，进行错误检查是非常重要的，
		// 即使在 defer 函数中也是如此。
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
