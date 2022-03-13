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
	PARSENUM = 5
	READ     = "./code.txt"
	RESULT   = "./result.go"
)

var CURField = map[string]*AutoStruct{}

// 1、需要生成字段名

type AutoStruct struct {
	Ref    string // 所属的结构体名字
	IsUse  bool   // 是否作为嵌套使用
	Cache  string //生成的结构体字符串缓存 多次使用
	fields []*AutoField
}

type AutoField struct {
	FieldName string
	FieldType string
	FieldDesc string
}

func main() {
	file, err := os.Open(READ)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	line := bufio.NewReader(file)
	// fields := []*AutoField{}

	var result bytes.Buffer
	autoStruct := new(AutoStruct)
	result.WriteString("package main\n")
	for {
		// content, _ := readLine(line, PARSENUM) // 随时调整策略
		content, _ := readWrods(line, PARSENUM) // 需要有这么多单词
		if content == "" {
			break
		}
		curLine := strings.TrimSpace(string(content))
		if strings.Index(curLine, "===") == 0 {
			autoStruct.Ref = curLine[3:]
			// result.WriteString(generateStruct(curLine[3:], fields))
			// result.WriteString(generateStructWithStruct(autoStruct))
			CURField[autoStruct.Ref] = autoStruct
			// fields = fields[:0]
			autoStruct = new(AutoStruct)
			// autoStruct.fields = autoStruct.fields[:0]
		} else if res := parseLine(string(content)); res != nil {
			// fields = append(fields, res)
			autoStruct.fields = append(autoStruct.fields, res)
		}
	}

	// 根据curfield生成字符串
	for _, s := range CURField {
		generateStructWithStruct(s)
	}
	for _, s := range CURField {
		if !s.IsUse {
			result.WriteString(s.Cache)
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

// 读取n个单词作为一行 不能包含空白符
func readWrods(r *bufio.Reader, n int) (string, error) {
	var res bytes.Buffer
	g := regexp.MustCompile(`[\S]+`)
	for {
		content, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if len(content) == 1 && content[0] == ' ' {
			res.WriteByte('-')
		} else {
			res.Write(content)
		}
		res.WriteByte(' ')

		// 如果以===开头表示一组解析已经解析完了
		if len(content) > 3 && string(content[0:3]) == "===" {
			break
		}

		// 检查是否有n个单词了
		if len(g.FindAll(res.Bytes(), -1)) >= n { // 因为存在最后的description有空格分割
			break
		}
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
			FieldDesc: strings.Join(parts[2:], " "),
		}
	default:
		return &AutoField{
			FieldName: parts[0],
			FieldType: parts[2],
			FieldDesc: strings.Join(parts[PARSENUM-1:], " "),
		}
	}
}

func generateStruct(structName string, fields []*AutoField) string {
	var byt bytes.Buffer
	byt.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, item := range fields {
		byt.WriteString(generateField(item.FieldName, item.FieldType, item.FieldDesc))
	}
	byt.WriteString("}\n")

	return byt.String()
}

// 根据AutoStruct生成
func generateStructWithStruct(auto *AutoStruct) string {
	if auto.Cache != "" {
		return auto.Cache
	}

	var byt bytes.Buffer
	byt.WriteString(fmt.Sprintf("type %s struct {\n", auto.Ref))
	for _, item := range auto.fields {
		byt.WriteString(fmt.Sprintf("%s %s %s %s\n",
			generateFieldName(item.FieldName),
			generateFieldTypeWithDeep(item.FieldType, auto.Ref),
			generateFieldTag(item.FieldName),
			generateFieldDesc(item.FieldDesc)))
	}
	byt.WriteString("}\n")

	auto.Cache = byt.String() // 缓存
	return auto.Cache
}

func generateField(name, t, desc string) string {
	return fmt.Sprintf("%s %s %s %s\n", generateFieldName(name), generateFieldType(t), generateFieldTag(name), generateFieldDesc(desc))
}

// 生成字段名 大写驼峰式
func generateFieldName(name string) string {
	return datautil.CamelString(name)
}

// 生成tag
func generateFieldTag(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}

// 生成注释
func generateFieldDesc(desc string) string {
	return "// " + desc
}

// 生成类型 嵌套struct
func generateFieldType(t string) string {
	switch strings.ToLower(t) {
	case "int":
		return "int"
	case "string", "text":
		return "string"
	case "float", "decimal":
		return "float64"
	case "datetime", "datatime", "date":
		return "time.Time"
	case "array":
		return "[]interface{}"
	case "object":
		return "interface{}"
	default:
		return t
	}
}

// 添加对嵌套类型的适配 通过structName 与 t的组合得出的字段名快速找到对应的结构体
func generateFieldTypeWithDeep(fieldTyp, structName string) string {
	if val, ok := CURField[structName+fieldTyp]; ok {
		generateStructWithStruct(val)
		val.IsUse = true // 如果匹配到发现这个类型对应的是一个结构体就把对应的isuse置为true
		return strings.TrimRight(val.Cache[5+len(structName+fieldTyp):], "\n") + " "
	}

	switch strings.ToLower(fieldTyp) {
	case "int":
		return "int"
	case "string", "text":
		return "string"
	case "float", "decimal":
		return "float64"
	case "datetime", "datatime", "date":
		return "time.Time"
	case "array":
		return "[]interface{}"
	case "object":
		return "interface{}"
	default:
		return fieldTyp
	}
}
