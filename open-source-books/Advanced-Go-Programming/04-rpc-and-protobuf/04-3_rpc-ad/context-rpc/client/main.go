package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":9090")
	if err != nil {
		log.Fatalf("error dialing tpc: %v", err)
	}

	err = client.Call("HelloService.Login", "user:pass", new(string))
	if err != nil {
		log.Fatalf("failed to log in: %v", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "kesa", &reply)
	if err != nil {
		log.Fatalf("error calling HelloService.Hello: %v", err)
	}

	fmt.Println(reply)
}
