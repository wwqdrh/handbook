package reflect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func StructGetTag() {
	type User struct {
		Name   string "user name"
		Passwd string "user password"
	}

	user := &User{"chronos", "pass"}

	s := reflect.TypeOf(user).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) // 将Tag打印出来
	}
}

////////////////////
// 动态创建struct 使用构造器 存储默认字段以及需要 增加的
// 由于静态语言 类型已经确定 所以这里只能新增struct而不能直接在原有基础上修改
// 然后通过这个数组中的来生成
////////////////////

func SturctNewTag() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
}

// 原始model 根据name动态修改

func StructModifyByName(fields2 ...reflect.StructField) {
	type User struct {
		Name float64 "user name"
	}

	user := &User{}

	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(float64(0)),
			Tag:  reflect.StructTag("new tag"),
		},
	}

	newFields := map[string]reflect.StructField{}

	elem := reflect.TypeOf(user).Elem()
	elemLen := elem.NumField()
	for i := 0; i < elemLen; i++ {
		field := elem.Field(i)
		newFields[field.Name] = field
	}

	for _, item := range fields {
		if _, ok := newFields[item.Name]; ok {
			newFields[item.Name] = item
		}
	}

	res := []reflect.StructField{}
	for _, v := range newFields {
		res = append(res, reflect.StructField{
			Name: v.Name,
			Type: v.Type,
			Tag:  v.Tag,
		})
	}

	fmt.Printf("%v\n", res)

	newUser := reflect.StructOf(res)

	elem = newUser.Elem()
	for i := 0; i < elem.NumField(); i++ {
		cur := elem.Field(i)
		fmt.Printf("%s %s", cur.Name, cur.Tag)
	}
}

// 遍历结构体字段
type User struct {
	Name string `json:"name" validator:"not_empty;in(1,2)"`
	Age  int    `json:"age" validator:"not_empty;in(1,2,3)"`
}

func IterStructField() {
	user := User{Name: "json", Age: 25}
	getType := reflect.TypeOf(user)
	getValue := reflect.ValueOf(user)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s %s: %v = %v\n", field.Name, field.Tag, field.Type, value)
	}
}
