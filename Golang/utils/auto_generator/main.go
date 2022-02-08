package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"wwqdrh/handbook/common/datautil"
)

const (
	PARSENUM = 1
	READ     = "./code.txt"
	RESULT   = "./result.go"
)

// 1、需要生成字段名
type AutoField struct {
	FieldName string
	FieldType string
}

func main() {
	file, err := os.Open(READ)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	line := bufio.NewReader(file)
	fields := []*AutoField{}

	var result bytes.Buffer
	result.WriteString("package main\n")
	for {
		content, _ := readLine(line, PARSENUM) // 随时调整策略
		if content == "" {
			break
		}
		curLine := strings.TrimSpace(string(content))
		if strings.Index(curLine, "===") == 0 {
			result.WriteString(generateStruct(curLine[3:], fields))
			fields = fields[:0]
		} else if res := parseLine(string(content)); res != nil {
			fields = append(fields, res)
		}
	}

	if err := ioutil.WriteFile(RESULT, result.Bytes(), 0644); err != nil {
		fmt.Println("===")
		fmt.Println(err)
		fmt.Println("===")
	}
}

// 读取n行作为一行
func readLine(r *bufio.Reader, n int) (string, error) {
	var res bytes.Buffer
	for i := 0; i < n; i++ {
		content, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		res.Write(content)
		res.WriteByte(' ')
	}
	return res.String(), nil
}

// 从一行的数据中解析field结构体
func parseLine(line string) *AutoField {
	// parts := strings.Fields(line) // 按空格分割
	// parts := strings.Split(line, "\t") // 按制表符分割
	// 使用re匹配
	r, _ := regexp.Compile(`\S+`)
	parts := r.FindAllString(line, -1)

	if len(parts) < 3 {
		return nil
	}

	switch len(parts) {
	case 3:
		return &AutoField{
			FieldName: parts[0],
			FieldType: parts[1],
		}
	default:
		return &AutoField{
			FieldName: parts[0],
			FieldType: parts[2],
		}
	}
}

func generateStruct(structName string, fields []*AutoField) string {
	var byt bytes.Buffer
	byt.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, item := range fields {
		byt.WriteString(generateField(item.FieldName, item.FieldType))
	}
	byt.WriteString("}\n")

	return byt.String()
}

func generateField(name, t string) string {
	return fmt.Sprintf("%s %s %s\n", generateFieldName(name), generateFieldType(t), generateFieldTag(name))
}

// 生成字段名 大写驼峰式
func generateFieldName(name string) string {
	return datautil.CamelString(name)
}

// 生成tag
func generateFieldTag(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}

// 生成类型
func generateFieldType(t string) string {
	switch strings.ToLower(t) {
	case "int":
		return "int"
	case "string":
		return "string"
	case "float", "decimal":
		return "float64"
	case "datetime", "datatime":
		return "time.Time"
	case "array", "object":
		return "[]string"
	default:
		return t
	}
}
