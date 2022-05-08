package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProxy(t *testing.T) {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
