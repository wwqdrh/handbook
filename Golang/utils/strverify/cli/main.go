package main

import (
	"context"
	"strverify"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(1 * time.Second)
	}()
	strverify.ServerStart(ctx, false)
}
