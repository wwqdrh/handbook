package shell

import (
	"crypto/rand"
	"encoding/hex"
)

func QuoteArg(s string) string {
	return "'" + strings.Replace(s, "'", "\"", -1) + "'"
}

func CreateRandomString(bytes int) string {
	c := bytes
	b := make([]byte, c)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}
