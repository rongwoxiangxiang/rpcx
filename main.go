package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"rpc/service"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Activity", new(service.ActivitService), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
