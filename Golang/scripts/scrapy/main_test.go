package main

import (
	"testing"
)

func compareList(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 测试
// func TestgetYearIDByRange(t *testing.T) {
func TestGetYearIDByRange(t *testing.T) {
	var res []string
	var expect []string

	res = getYearIDByRange(1960, 1960)
	expect = []string{"1"}
	if !compareList(res, expect) {
		t.Error("数据失败")
	}

	res = getYearIDByRange(1960, 2010)
	expect = []string{"1", "2", "3", "4", "5"}
	if !compareList(res, expect) {
		t.Error("数据失败")
	}

	res = getYearIDByRange(1970, 2012)
	expect = []string{"2", "3", "4", "5", "6", "7"}
	if !compareList(res, expect) {
		t.Error("数据失败")
	}

	res = getYearIDByRange(1960, 2022)
	expect = []string{"0"}
	if !compareList(res, expect) {
		t.Error("数据失败")
	}
}

func TestSearchMovieIDByYearID(t *testing.T) {
	searchMovieIDByYearID("6")
}
