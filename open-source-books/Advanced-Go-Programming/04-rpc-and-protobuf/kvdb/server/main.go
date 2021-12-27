package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-and-protobuf/kvdb/service"
)

func main() {
	err := RegisterKVStoreService()
	if err != nil {
		log.Fatalf("error registering %s: %v", service.KVStoreName, err)
	}

	log.Printf("listening %s ...", ":9090")

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("error listening tcp: %v", err)
	}

	log.Println("waiting for connection ...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accept conn on %s: %v", conn.RemoteAddr(), err)
		}
		go rpc.ServeConn(conn)
	}
}

func RegisterKVStoreService() error {
	return rpc.RegisterName(service.KVStoreName, service.NewKVStoreService())
}
