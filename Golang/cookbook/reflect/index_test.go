package reflect

import (
	"reflect"
	"testing"
)

func TestString2Slice(t *testing.T) {
	if !reflect.DeepEqual(String2Slice("aaa"), []byte{'a', 'a', 'a'}) {
		t.Error("发生错误")
	}
}
