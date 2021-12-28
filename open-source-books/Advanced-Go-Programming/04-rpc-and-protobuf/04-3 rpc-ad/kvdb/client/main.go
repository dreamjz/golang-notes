package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	service2 "rpc-and-protobuf/04-3 rpc-ad/kvdb/service"
)

const (
	timeout = 10
)

type KVStoreClient struct {
	client *rpc.Client
}

func (kvs *KVStoreClient) Get(key string, value *string) error {
	return kvs.client.Call(service2.KVStoreName+".Get", key, value)
}

func (kvs *KVStoreClient) Set(kv []string, reply *struct{}) error {
	return kvs.client.Call(service2.KVStoreName+".Set", kv, reply)
}

func (kvs *KVStoreClient) Watch(timeoutSecond int, keyChanged *string) error {
	return kvs.client.Call(service2.KVStoreName+".Watch", timeoutSecond, keyChanged)
}

var _ service2.KVStoreInterface = (*KVStoreClient)(nil)

func DialKVStoreService(network string, addr string) (*KVStoreClient, error) {
	client, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	kvClient := &KVStoreClient{
		client: client,
	}
	return kvClient, nil
}

var (
	cmd    string
	params []string
)

func main() {
	parse()

	client, err := DialKVStoreService("tcp", ":9090")
	if err != nil {
		log.Fatalf("error dialing tcp: %v", err)
	}

	switch cmd {
	case "get":
		var val string
		err := get(client, params[0], &val)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
	case "set":
		err := set(client, params)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unsupported command")
	}
}

func parse() {
	if len(os.Args) < 3 {
		fmt.Print("USAGE: \n\t client [command] [...params]")
		os.Exit(1)
	}
	cmd = os.Args[1]
	params = os.Args[2:]
	log.Printf("Cmd: %s, params: %v", cmd, params)
}

func doClientWork(client *KVStoreClient, srcMethod func() error, watched bool) error {
	if !watched {
		return srcMethod()
	}
	watchChan := make(chan bool)
	go watchKeyChanged(client, watchChan)
	err := srcMethod()
	<-watchChan
	return err
}

func watchKeyChanged(client *KVStoreClient, watchChan chan bool) {
	var keyChanged string
	err := client.Watch(timeout, &keyChanged)
	if err != nil {
		log.Printf("error calling %s: %v", service2.KVStoreName+".Watch", err)
	}
	log.Print("watch-key-changed: ", keyChanged)
	watchChan <- true
}

func get(client *KVStoreClient, key string, val *string) error {
	return doClientWork(client, func() error {
		err := client.Get(key, val)
		if err != nil {
			return fmt.Errorf("error calling %s.%s: %w", service2.KVStoreName, "Get", err)
		}
		return nil
	}, false)
}

func set(client *KVStoreClient, kv []string) error {
	return doClientWork(client, func() error {
		err := client.Set(kv, new(struct{}))
		if err != nil {
			return fmt.Errorf("error calling %s.%s: %v", service2.KVStoreName, "Set", err)
		}
		return nil
	}, true)
}
