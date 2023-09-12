package cookbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectMap(t *testing.T) {
	reflectMap()
}

func TestReflectSlice(t *testing.T) {
	reflectSlice()
}

func TestStringBytes(t *testing.T) {
	word := "hello world"

	wordByte := string2bytes(word)
	assert.Equal(t, word, bytes2string(wordByte))
}
