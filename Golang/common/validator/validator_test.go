package validator

import (
	"fmt"
	"testing"
)

type PageInfo struct {
	Page     int `json:"page" form:"page" validator:"not_empty;in(1,2,3)"` // 页码
	PageSize int `json:"pageSize" form:"pageSize" validator:"not_empty"`   // 每页大小
}

// 1、读取结构体中的有哪些字段，配置了什么校验类型，值是多少 字段名字(如果有json格式的就用这个更好展示)
func TestIterStructField(t *testing.T) {
	IterStructField(PageInfo{Page: 1, PageSize: 1})
	IterStructField(&PageInfo{Page: 1, PageSize: 1})
}

// 获取校验类型 进行校验 传入FieldMeta
func TestFieldValidate(t *testing.T) {
	for _, item := range IterStructField(PageInfo{}) {
		fmt.Println(ValidatorField(item))
	}
}
