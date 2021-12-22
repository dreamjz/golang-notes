package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-and-protobuf/json-rpc/rpc_objects"
)

func main() {
	calc := new(rpc_objects.Args)
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
