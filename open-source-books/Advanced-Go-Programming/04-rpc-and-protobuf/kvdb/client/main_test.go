package main

import (
	"fmt"
	"testing"
	"time"
)

const num = 40000

func TestSet(t *testing.T) {
	client, err := DialKVStoreService("tcp", ":9090")
	if err != nil {
		t.Fatalf("error dialing tcp: %v", err)
	}
	retChan := make(chan error, num)
	for i := 0; i < 2; i++ {
		go func() {
			for i := 20000; i < num; i++ {
				key := fmt.Sprintf("k-%d", i)
				val := fmt.Sprintf("v-%d", i)
				kv := []string{key, val}
				go func() {
					time.Sleep(5 * time.Second)
					err := set(client, kv)
					retChan <- err
				}()
			}
		}()
	}
	//for i := 0; i < 2*20000; i++ {
	//	if err := <-retChan; err != nil {
	//		fmt.Println(err)
	//	}
	//}
	for err := range retChan {
		if err != nil {
			fmt.Println(err)
		}
	}
}
