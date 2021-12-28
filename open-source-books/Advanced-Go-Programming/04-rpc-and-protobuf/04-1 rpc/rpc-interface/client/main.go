package main

import (
	"fmt"
	"log"
	"net/rpc"
	rpc_objects2 "rpc-and-protobuf/04-1 rpc/rpc-interface/rpc_objects"
)

const HelloServiceName = "path/example/HelloService"

var _ rpc_objects2.HelloInterface = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

func DialHelloService(network string, addr string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		Client: client,
	}, nil
}

func main() {
	client, err := DialHelloService("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to dial tcp: %v", err)
	}
	var reply string
	err = client.Hello("kesa", &reply)
	if err != nil {
		log.Fatalf("call Hello error: %v", err)
	}
	fmt.Println("Result: ", reply)
}
