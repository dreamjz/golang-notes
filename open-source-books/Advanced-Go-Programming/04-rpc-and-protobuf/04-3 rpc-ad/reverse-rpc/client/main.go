package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	service2 "rpc-and-protobuf/04-3 rpc-ad/reverse-rpc/service"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("error listening tcp:", err)
	}
	log.Print("listening on tcp :9090")
	clientChan := make(chan *rpc.Client)

	go func() {
		log.Println("waiting for connection ...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("error accepting conn:", err)
		}
		log.Printf("%s connected ...", conn.RemoteAddr())
		clientChan <- rpc.NewClient(conn)
	}()

	doClientWork(clientChan)
}

func doClientWork(clientChan chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()
	var reply int
	args := service2.ArithmeticService{
		X: 2,
		Y: 3,
	}
	err := client.Call("ArithmeticService.Multiply", args, &reply)
	if err != nil {
		log.Println("error calling ArithmeticService.Multiply:", err)
	}

	fmt.Printf("%d * %d = %d", args.X, args.Y, reply)
}
