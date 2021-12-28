package main

import (
	"log"
	"net"
	"net/rpc"
	rpc_objects2 "rpc-and-protobuf/04-1 rpc/rpc-interface/rpc_objects"
)

const HelloServiceName = "path/example/HelloService"

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func RegisterHelloService(svc rpc_objects2.HelloInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen tcp: %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept conn: %v", err)
		}
		go rpc.ServeConn(conn)
	}
}
