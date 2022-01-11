package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpc_objects2 "rpc-and-protobuf/04-1_rpc/json-rpc/rpc_objects"
)

func main() {
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to dial tcp: %v", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	args := &rpc_objects2.Args{M: 2, N: 3}
	var reply int

	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatalf("failed to call rpc: %v", err)
	}

	fmt.Printf("%d * %d = %d", 2, 3, reply)
}
