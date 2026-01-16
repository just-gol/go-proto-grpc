package main

import (
	"context"
	"fmt"
	"greeter/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

type Hello struct {
	greeter.UnimplementedGreeterServer
}

func (h Hello) SayHello(ctx context.Context, req *greeter.HelloReq) (*greeter.HelloResp, error) {
	fmt.Println(req)
	return &greeter.HelloResp{
		Message: "Hello " + req.Name,
	}, nil

}

func main() {
	server := grpc.NewServer()
	greeter.RegisterGreeterServer(server, new(Hello))
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
	}
	defer listen.Close()
	server.Serve(listen)
}
