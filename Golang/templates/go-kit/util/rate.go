package util

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"fmt"
	"time"
)

// 限流，漏桶、令牌桶
func rateSample() {
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()
	for {
		err := r.WaitN(ctx, 2) // wait
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().Format("2020-12-31 23:59:59"))
		time.Sleep(time.Second)
	}
}

