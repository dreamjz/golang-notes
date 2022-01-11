package main

import (
	"context"
	"flag"
	"log"
	pb "rpc-and-protobuf/04-4_gRPC/hello"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:9090", "The server address in format of host:port")
	name       = flag.String("name", defaultName, "Name to greet")
)

func sayHello(client pb.GreeterClient, name *pb.HelloRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	message, err := client.SayHello(ctx, name)
	if err != nil {
		log.Fatalf("%v.SayHello(_) = _, %v", client, err)
	}
	log.Println(message)
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect %v: %v", *serverAddr, err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	sayHello(client, &pb.HelloRequest{Name: *name})
}
