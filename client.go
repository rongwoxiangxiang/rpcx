package main

import (
	"./lib/service"
	pb "./lib/service" // 引入proto包
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// Address gRPC服务地址
	Address2 = "127.0.0.1:50052"
)

func main() {
	log.Println(">>>>>>>>")
	// 连接
	conn, err := grpc.Dial(Address2, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := pb.NewActivityServiceClient(conn)
	// 调用方法
	//reqBody := new(pb.Checkin)
	//reqBody.Wid = 211
	//reqBody.ActivityId = 211
	//reqBody.Total = 111
	//reqBody.Lastcheckin = time.Now().Unix()

	reqBody := new(service.ActivityQuery)
	reqBody.Wid = 1
	reqBody.Index = 1
	reqBody.Limit = 10

	t1 := time.Now().Unix()
	r, err := c.List(context.Background(), reqBody)

	if err != nil {
		log.Fatalln(err)
	}
	t2 := time.Now().Unix()
	log.Println(t2 - t1)
	log.Println(r)
}
