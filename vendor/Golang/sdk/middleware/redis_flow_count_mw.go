package middleware

import (
	"fmt"
	"wwqdrh/handbook/sdk/middleware/public"
)

func RedisFlowCountMiddleWare(counter *public.RedisFlowCountService) func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		counter.Increase()
		fmt.Println("QPS:", counter.QPS)
		fmt.Println("TotalCount:", counter.TotalCount)
		c.Next()
	}
}
