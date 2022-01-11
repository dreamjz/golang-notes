package main

import (
	"log"
	"net"
	"net/rpc"
	service2 "rpc-and-protobuf/04-3_rpc-ad/context-rpc/service"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("error listening tcp:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error accepting conn:", err)
		}
		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&service2.HelloService{Conn: conn})
			p.ServeConn(conn)
		}()
	}
}
