package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type validateType int

const (
	NotNull validateType = iota
)

type FieldMeta struct {
	Name     string            // 字段名
	NickName string            // 可能由json等指定
	Tag      map[string]string // 标签 按照空格分割
	Value    reflect.Value     // 字段值
	Type     reflect.Type      // 字段的类型
}

func Validator(data interface{}) error {
	for _, item := range IterStructField(data) {
		if err := ValidatorField(item); err != nil {
			return err
		}
	}
	return nil
}

// 返回结构体的所有字段、值、tag
func IterStructField(data interface{}) []*FieldMeta {
	res := []*FieldMeta{}
	getType := reflect.TypeOf(data)
	getValue := reflect.ValueOf(data)

	if fmt.Sprint(getType.Kind()) == "ptr" {
		getType = getType.Elem()
		getValue = getValue.Elem()
	}

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		// value := getValue.Field(i).Interface()

		tagMap := map[string]string{}
		for _, item := range strings.Fields(string(field.Tag)) {
			t := strings.Split(item, ":")
			tagMap[t[0]] = t[1][1 : len(t[1])-1] // 将"“去掉
		}
		nickName := field.Name
		if val, ok := tagMap["json"]; ok {
			nickName = val
		}

		res = append(res, &FieldMeta{
			Name:     field.Name,
			NickName: nickName,
			Value:    getValue.Field(i),
			Tag:      tagMap,
			Type:     field.Type,
		})
	}
	return res
}

func ValidatorField(field *FieldMeta) error {
	var (
		val string
		ok  bool
	)
	if val, ok = field.Tag["validator"]; !ok {
		return nil
	}
	errMsg := []string{}
	for _, item := range strings.Split(val, ";") {
		if err := validator(item, field.NickName, field.Value); err != nil {
			errMsg = append(errMsg, err.Error())
		}
	}
	if len(errMsg) == 0 {
		return nil
	}
	return errors.New(strings.Join(errMsg, "\n"))
}

// 校验 not_empty
func validator(validate string, name string, value reflect.Value) error {
	switch {
	case validate == "not_empty":
		return notEmpty(name, value)
	case strings.Index(validate, "in") == 0:
		return in(name, value, validate)
	default:
		return nil
	}
}

func notEmpty(name string, value reflect.Value) error {
	var flag bool
	switch value.Kind() {
	case reflect.String:
		flag = value.Len() == 0
	case reflect.Bool:
		flag = !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		flag = value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		flag = value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		flag = value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		flag = value.IsNil()
	default:
		flag = reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
	if flag {
		return fmt.Errorf("%s: 不能为空", name)
	} else {
		return nil
	}
}

func in(name string, value reflect.Value, validate string) error {
	validate = validate[3 : len(validate)-1]
	flag := false
	for _, item := range strings.Split(validate, ",") {
		if item == fmt.Sprint(value.Interface()) {
			flag = true
			break
		}
	}
	if !flag {
		return fmt.Errorf("%s: 需要 %s", name, validate)
	}
	return nil
}

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RegisterRule
//@description: 注册自定义规则方案建议在路由初始化层即注册
//@param: key string, rule Rules
//@return: err error

func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + "已注册,无法重复注册")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: NotEmpty
//@description: 非空 不能为其对应类型的0值
//@return: string

func NotEmpty() string {
	return "notEmpty"
}

//@author: [zooqkl](https://github.com/zooqkl)
//@function: RegexpMatch
//@description: 正则校验 校验输入项是否满足正则表达式
//@param:  rule string
//@return: string
func RegexpMatch(rule string) string {
	return "regexp=" + rule
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Lt
//@description: 小于入参(<) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Lt(mark string) string {
	return "lt=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Le
//@description: 小于等于入参(<=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Le(mark string) string {
	return "le=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Eq
//@description: 等于入参(==) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Eq(mark string) string {
	return "eq=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ne
//@description: 不等于入参(!=)  如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Ne(mark string) string {
	return "ne=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ge
//@description: 大于等于入参(>=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Ge(mark string) string {
	return "ge=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Gt
//@description: 大于入参(>) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
//@param: mark string
//@return: string

func Gt(mark string) string {
	return "gt=" + mark
}

// In 判断值是否是其中一个 用逗号隔开
func In(mark string) string {
	return "in=" + mark
}

//
//@author: [piexlmax](https://github.com/piexlmax)
//@function: Verify
//@description: 校验方法
//@param: st interface{}, roleMap Rules(入参实例，规则map)
//@return: err error

func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
		"in": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	kd := val.Kind() // 获取到st对应的类别
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				case strings.Split(v, "=")[0] == "regexp":
					if !regexpMatch(strings.Split(v, "=")[1], val.String()) {
						return errors.New(tagVal.Name + "格式校验不通过")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "长度或值不在合法范围," + v)
					}
				}
			}
		}
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: compareVerify
//@description: 长度和数字的校验方法 根据类型自动校验
//@param: value reflect.Value, VerifyStr string
//@return: bool

func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: isBlank
//@description: 非空校验
//@param: value reflect.Value
//@return: bool

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: compare
//@description: 比较函数
//@param: value interface{}, VerifyStr string
//@return: bool

func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)

	// in操作
	if VerifyStrArr[0] == "in" {
		flag := false
		for _, item := range strings.Split(VerifyStrArr[1], ",") {
			if item == fmt.Sprint(val) {
				flag = true
				break
			}
		}
		return flag
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		case VerifyStrArr[0] == "in":
			flag := false
			for _, item := range strings.Split(VerifyStrArr[1], ",") {
				if item == val.String() {
					flag = true
					break
				}
			}
			return flag
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}
