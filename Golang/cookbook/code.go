package cookbook

import (
	"strconv"
	"strings"
)

// 编码相关
func Zh2unicode(word string) string {
	textQuoted := strconv.QuoteToASCII(word)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}

func Unicode2zh(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
