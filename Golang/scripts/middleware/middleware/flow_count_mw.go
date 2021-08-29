package middleware

import (
	"fmt"
	"wwqdrh/handbook/scripts/middleware/public"
)

func FlowCountMiddleWare(counter *public.FlowCountService) func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		counter.Increase()
		fmt.Println("QPS:", counter.QPS)
		fmt.Println("TotalCount:", counter.TotalCount)
		c.Next()
	}
}
