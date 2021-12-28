package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpc_objects2 "rpc-and-protobuf/04-1 rpc/json-rpc/rpc_objects"
)

func main() {
	calc := new(rpc_objects2.Args)
	rpc.Register(calc)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen tcp: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept conn: %v", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
