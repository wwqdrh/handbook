package main

import (
	"context"
	"fmt"

	"wwqdrh/handbook/tools/rpc/grpc/greeter"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greeter.NewGreeterServiceClient(conn)

	ctx := context.Background()
	req := greeter.GreetRequest{Greeting: "Hello", Name: "Reader"}
	resp, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	req.Greeting = "Goodbye"
	resp, err = client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
