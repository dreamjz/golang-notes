package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("create new watcher failed:", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s : %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("../resources")
	if err != nil {
		log.Fatal("add watch path failed:", err)
	}
	<-done
}
