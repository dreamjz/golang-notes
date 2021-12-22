package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatalf("rpc error: %v", err)
	}

	fmt.Println("Result: ", reply)
}
