package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件需要经常进行错误检查，
// 这个帮助方法可以精简下面的错误检查过程。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileInExample() {

	// 最基本的文件读取任务或许就是将文件内容读取到内存中。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 您通常会希望对文件的读取方式和内容进行更多控制。
	// 对于这个任务，首先使用 `Open` 打开一个文件，以获取一个 `os.File` 值。
	f, err := os.Open("/tmp/dat")
	check(err)

	// 从文件的开始位置读取一些字节。
	// 最多允许读取 5 个字节，但还要注意实际读取了多少个。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 你也可以 `Seek` 到一个文件中已知的位置，并从这个位置开始读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 例如，`io` 包提供了一个更健壮的实现 `ReadAtLeast`，用于读取上面那种文件。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内建的倒带，但是 `Seek(0, 0)` 实现了这一功能。
	_, err = f.Seek(0, 0)
	check(err)

	//  `bufio` 包实现了一个缓冲读取器，这可能有助于提高许多小读操作的效率，以及它提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 任务结束后要关闭这个文件
	// （通常这个操作应该在 `Open` 操作后立即使用 `defer` 来完成）。
	f.Close()
}

func ReadFile() {
	//1、一次性读取文件内容,还有一个 ReadAll的函数，也能读取
	data, err := ioutil.ReadFile("./util/file.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	//2、逐行读取
	file, err := os.Open("./util/file.go") //打开
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //关闭

	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(content))
	}

	//3、按照字节数读取
	file, err = os.Open("./util/file.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//读取数据
	bs := make([]byte, 4)
	for {
		_, err = file.Read(bs)
		if err == io.EOF {
			break
		}
		fmt.Print(string(bs))
	}
}

func WriteFile() {
	// 会覆盖 不存在则创建
	content := []byte("测试1\n测试2\n")
	err := ioutil.WriteFile("test.txt", content, 0644)
	if err != nil {
		panic(err)
	}

	// 添加新内容
	var str = "测试1\n测试2\n"
	var filename = "./test.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	if err1 != nil {
		panic(err1)
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, _ := w.WriteString(str)
	fmt.Printf("写入 %d 个字节n", n)
	w.Flush()
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
