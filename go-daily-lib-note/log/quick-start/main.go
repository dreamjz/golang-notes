package main

import "log"

type User struct {
	ID       int
	Username string
}

func main() {
	user := User{
		ID:       1,
		Username: "Kesa",
	}

	// set prefix
	log.SetPrefix("[Login]")

	log.Printf("User: %s login,ID:%d", user.Username, user.ID)
	log.Fatalf("Warning hacker %s detected", user.Username)

	// do not execute
	log.Panicf("User:%s[ID:%d] login failed", user.Username, user.ID)

	// different behaviors between log.Print and log.Println
	log.Print("A", 1, 2, "B")
	log.Println("A", 1, 2, "B")
}
