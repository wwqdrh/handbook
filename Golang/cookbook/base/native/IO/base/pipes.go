package main

import (
	"io"
	"os"
)

func PipeExample() error {
	r, w := io.Pipe()
	go func() {
		_, _ = w.Write([]byte("test\n"))
		_ = w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
