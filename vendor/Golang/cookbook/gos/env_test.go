package gos

import (
	"fmt"
	"testing"
)

func TestGetEnv(t *testing.T) {
	fmt.Println(GetEnv("LOG_LEVEL"))
}
