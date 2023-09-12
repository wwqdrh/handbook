package main

import (
	"log"
	"net/http"
	"wwqdrh/handbook/cmd/middleware/load_balance"
	"wwqdrh/handbook/cmd/middleware/middleware"
	"wwqdrh/handbook/cmd/middleware/proxy"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("http://127.0.0.1:2003", "50")
	proxy := proxy.NewLoadBalanceReverseProxy(&middleware.SliceRouterContext{}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
