package main

import "go-jwt-note/initialize"

func main() {
	initialize.Redis()
	initialize.Run()
}
