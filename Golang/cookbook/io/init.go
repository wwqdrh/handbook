package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// PipeExample helps give some more examples of using io interfaces
func PipeExample() error {
	// the pipe reader and pipe writer implement
	// io.Reader and io.Writer
	r, w := io.Pipe()

	// this needs to be run in a separate go routine
	// as it will block waiting for the reader
	// close at the end for cleanup
	go func() {
		// for now we'll write something basic,
		// this could also be used to encode json
		// base64 encode, etc.
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}

// Copy copies data from in to out first directly,
// then using a buffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// we write to out, but also Stdout
	w := io.MultiWriter(out, os.Stdout)

	// a standard copy, this can be dangerous if there's a lot
	// of data in in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	// buffered write using 64 byte chunks
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// lets print a new line
	fmt.Println()

	return nil
}

func CopyExample() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample = ")
	if err := PipeExample(); err != nil {
		panic(err)
	}
}
