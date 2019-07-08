package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc/lib/config"
	"rpc/lib/service"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	config.InitConfig()
	config.InitStoreDb()
	listen, err := net.Listen(
		"tcp",
		config.Conf.GetDefault("application.address", "127.0.0.1:50052").(string),
	)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc Server
	s := grpc.NewServer()
	service.RegisterActivityServiceServer(s, service.ActivityServiceImpl{})
	service.RegisterCheckinServiceServer(s, service.CheckinServiceImpl{})
	log.Println("Listen on " + Address)
	s.Serve(listen)
}
