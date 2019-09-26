package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"log"
	"rpc/common"
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

	args := &pb.RequestList{Limit: 20, Params: map[string]string{"wid": "2", "used": common.NO_VALUE_STRING}}
	reply := &pb.PrizeList{}
	err := xclient.Call(context.Background(), "LimitByActivityIdAndUsed", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply)

}
