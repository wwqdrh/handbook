package main

import (
	"errors"
	"io"
	"os"
)

func Operate() error {
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	value := []byte("hello\n")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.Open("test.txt")
	if err != nil {
		return nil
	}
	_, _ = io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		return err
	}

	// 跳转到 /tmp 文件夹
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// 删除建立的文件夹
	// os.RemoveAll如果传递了错误的文件夹路径会返回错误
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}
