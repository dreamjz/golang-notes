package main

import (
	"log"
	"net"
	"net/rpc"
	service2 "rpc-and-protobuf/04-3_rpc-ad/reverse-rpc/service"
	"time"
)

func main() {
	rpc.Register(service2.NewArithmeticService())

	for {
		conn, err := net.Dial("tcp", ":9090")
		if err != nil {
			log.Printf("error dialing tcp: %v, retry ...", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("dialing success")

		rpc.ServeConn(conn)
		conn.Close()
	}
}
