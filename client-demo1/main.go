package main

import (
	"client-demo1/proto/greeter"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	}
	client := greeter.NewGreeterClient(conn)
	hello, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "hello",
	})
	if err != nil {
	}
	fmt.Printf("%+v\n", hello)
}
