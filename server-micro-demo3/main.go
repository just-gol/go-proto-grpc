package main

import (
	"goods/handler"
	pb "goods/proto"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
	"go-micro.dev/v5/server"
)

var (
	name    = "goods"
	version = "0.0.1"
)

func main() {
	registry := consul.NewRegistry()
	// Create service
	service := micro.NewService(
		micro.Name(name),       // ① 服务名字（别人发现服务靠它）
		micro.Version(version), // ② 版本号（可选，用于区分版本）
		micro.Server(server.NewServer(
			server.Name(name),                  // ① 同样设置服务名字（因为你自己创建了 Server）registry 时用的名字
			server.Registry(registry),          // ③ 用哪个注册中心（Consul/Nacos 的客户端实例）
			server.Address("0.0.0.0:9098"),     // ④ 本机监听地址：真正开端口的地方
			server.Advertise("127.0.0.1:9098"), // ⑤ 注册到注册中心的地址：告诉别人怎么连你
		)),
	)

	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterGoodsHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
