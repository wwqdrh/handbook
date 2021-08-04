package main

import (
	"fmt"
	"io"
	"os"
)

/**
使用buffer作为缓存将数据从in复制到out
*/
func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout) // 将结果写入out以及Stdout
	if _, err := io.Copy(w, in); err != nil {
		return nil
	}

	_, _ = in.Seek(0, 0)
	buf := make([]byte, 64) // 创建缓冲池
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}
	fmt.Println()
	return nil
}
