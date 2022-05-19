package system

import (
	"fmt"
	"testing"
)

func TestWalkPath(t *testing.T) {
	res, err := pathWalk("./", "_test.go")
	if err != nil {
		t.Error(err)
	} else {
		for _, item := range res {
			fmt.Println(item.Name())
		}
	}
}

func TestPathWalkDir(t *testing.T) {
	res, err := FilePathWalkDir("./", "_test.go")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestIOReadDir(t *testing.T) {
	res, err := IOReadDir("./")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestOSReadDir(t *testing.T) {
	res, err := OSReadDir("./")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}
