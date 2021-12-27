package service

import (
	"errors"
	"log"
	"sync"
	"time"
)

const (
	KVStoreName = "KVStoreService"
)

var (
	ErrNotFound = errors.New("not found")
	ErrTimeout  = errors.New("timed out")
)

type KVStoreInterface interface {
	Get(key string, value *string) error
	Set(kv []string, reply *struct{}) error
	Watch(timeoutSecond int, keyChanged *string) error
}

type KVStoreService struct {
	m         map[string]string
	filter    map[string]func(key string)
	watchChan chan string
	mu        sync.Mutex
}

var _ KVStoreInterface = (*KVStoreService)(nil)

func NewKVStoreService() *KVStoreService {
	kvs := &KVStoreService{
		m:         make(map[string]string),
		filter:    make(map[string]func(key string)),
		watchChan: make(chan string, 10),
	}
	kvs.filter["watcher"] = func(key string) {
		kvs.watchChan <- key
	}
	return kvs
}

func (kvs *KVStoreService) Get(key string, value *string) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	if v, ok := kvs.m[key]; ok {
		*value = v
		return nil
	}

	return ErrNotFound
}

func (kvs *KVStoreService) Set(kv []string, reply *struct{}) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	key, value := kv[0], kv[1]

	oldValue := kvs.m[key]
	log.Printf("old: %v, new: %v", oldValue, value)
	for _, fn := range kvs.filter {
		fn(key)
	}

	kvs.m[key] = value
	return nil
}

func (kvs *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return ErrTimeout
	case key := <-kvs.watchChan:
		*keyChanged = key
		return nil
	}
}
