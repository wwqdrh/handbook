package cookbook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZh2unicode(t *testing.T) {
	unicodeStr := Zh2unicode("你好, 世界")
	v, _ := Unicode2zh([]byte(unicodeStr))

	require.Equal(t, string(v), "你好, 世界")
}
