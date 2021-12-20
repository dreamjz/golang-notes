package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("listen tcp error: %v", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("accept error: %v", err)
	}

	rpc.ServeConn(conn)
}
