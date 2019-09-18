package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"log"
	"rpc/pb"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("Checkin", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := &pb.RequestById{Id: 2}
	reply := &pb.ResponseEffect{}
	err := xclient.Call(context.Background(), "LimitByWid", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply)

}
