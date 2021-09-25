package io

import (
	"bytes"
	"encoding/base64"
	b64 "encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

// Base64Example demonstrates using
// the base64 package
func Base64Example() error {
	// base64 is useful for cases where
	// you can't support binary formats
	// it operates on bytes/strings

	// using helper functions and URL encoding
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodeToString and URLEncoding: ", value)

	// decode the first value
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeToString and URLEncoding: ", string(decoded))

	return nil
}

// Base64ExampleEncoder shows similar examples
// with encoders/decoders
func Base64ExampleEncoder() error {
	// using encoder/ decoder
	buffer := bytes.Buffer{}

	// encode into the buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)

	// be sure to close
	if err := encoder.Close(); err != nil {
		return err
	}
	if _, err := encoder.Write([]byte("encoding some other data")); err != nil {
		return err
	}

	fmt.Println("Using encoder and StdEncoding: ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}

	fmt.Println("Using decoder and StdEncoding: ", string(results))

	return nil
}

// pos stores the x, y position
// for Object
type pos struct {
	X      int
	Y      int
	Object string
}

// GobExample demonstrates using
// the gob package
func GobExample() error {
	buffer := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	// note that if p was an interface
	// we'd have to call gob.Register first

	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	// note this is a binary format so it wont print well
	fmt.Println("Gob Encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil
}

// Go 提供了对 [base64 编解码](http://zh.wikipedia.org/wiki/Base64)的内建支持。

// 这个语法引入了 `encoding/base64` 包，
// 并使用别名 `b64` 代替默认的 `base64`。这样可以节省点空间。

func EncodingExample() {

	// 这是要编解码的字符串。
	data := "abc123!?$*&()'-=@~"

	// Go 同时支持标准 base64 以及 URL 兼容 base64。
	// 这是使用标准编码器进行编码的方法。
	// 编码器需要一个 `[]byte`，因此我们将 string 转换为该类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确，
	// 那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 URL base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
