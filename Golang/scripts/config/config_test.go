package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	// assert.Equal(t, Conf.Server.Addr, ":8090")
	fmt.Printf(Conf.Server.Addr)
}
