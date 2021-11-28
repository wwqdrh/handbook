package gof

import (
	"sync"
)

type SingletonType struct {
}

var instance *SingletonType
var once sync.Once

func GetInstance() *SingletonType {
	once.Do(func() {
		instance = &SingletonType{}
	})
	return instance
}
