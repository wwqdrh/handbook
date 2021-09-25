package os

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFileExample() {

	// 开始！这里展示了如何写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 对于更细粒度的写入，先打开一个文件。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// 打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 `Close`。
	defer f.Close()

	// 您可以按期望的那样 `Write` 字节切片。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString` 也是可用的。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 `Sync` 将缓冲区的数据写入硬盘。
	f.Sync()

	// 与我们前面看到的带缓冲的 Reader 一样，`bufio` 还提供了的带缓冲的 Writer。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用 `Flush` 来确保，已将所有的缓冲操作应用于底层 writer。
	w.Flush()

}

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

// Capitalizer opens a file, reads the contents,
// then writes those contents to a second file
func Capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, 0); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)

	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}

	s := strings.ToUpper(tmp.String())

	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	return nil
}

// CapitalizerExample creates two files, writes to one
//then calls Capitalizer() on both
func CapitalizerExample() error {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}

	if _, err := f1.Write([]byte(`
    this file contains
    a number of words
    and new lines`)); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := Capitalizer(f1, f2); err != nil {
		return err
	}

	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {
	// If you need a temporary place to store files with the
	// same name ie. template1-10.html a temp directory is a good
	// way to approach it, the first argument being blank means it
	// will use create the directory in the location returned by
	// os.TempDir()
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	// This will delete everything inside the temp file when this
	// function exits if you want to do this later, be sure to return
	// the directory name to the calling function
	defer os.RemoveAll(t)

	// the directory must exist to create the tempfile
	// created. t is an *os.File object.
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	// normally we'd delete the temporary file here, but because
	// we're placing it in a temp directory, it gets cleaned up
	// by the earlier defer

	return nil
}

func TempFileExample() {

	// 创建临时文件最简单的方法是调用 `ioutil.TempFile` 函数。
	// 它会创建并打开文件，我们可以对文件进行读写。
	// 函数的第一个参数传 `""`，`ioutil.TempFile` 会在操作系统的默认位置下创建该文件。
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// 打印临时文件的名称。
	// 文件名以 `ioutil.TempFile` 函数的第二个参数作为前缀，
	// 剩余的部分会自动生成，以确保并发调用时，生成不重复的文件名。
	// 在类 Unix 操作系统下，临时目录一般是 `/tmp`。
	fmt.Println("Temp file name:", f.Name())

	// defer 删除该文件。
	// 尽管操作系统会自动在某个时间清理临时文件，但手动清理是一个好习惯。
	defer os.Remove(f.Name())

	// 我们可以向文件写入一些数据。
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 如果需要写入多个临时文件，最好是为其创建一个临时 *目录* 。
	// `ioutil.TempDir` 的参数与 `TempFile` 相同，
	// 但是它返回的是一个 *目录名* ，而不是一个打开的文件。
	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// 现在，我们可以通过拼接临时目录和临时文件合成完整的临时文件路径，并写入数据。
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
