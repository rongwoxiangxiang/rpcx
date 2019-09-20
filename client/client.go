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

	discovery := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("Prize", client.Failtry, client.RandomSelect, discovery, opt)
	defer xclient.Close()

	args := &pb.PrizeAdd{Wid: 1, ActivityId: 2, Codes: []string{"111111"}}
	reply := &pb.ResponseEffect{}
	err := xclient.Call(context.Background(), "InsertBatch", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply)

}
