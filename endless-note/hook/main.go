package main

import (
	"log"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func preSigUsr1() {
	log.Println("pre SIGUSR1")
}

func postSigUsr1() {
	log.Println("pre SIGUSR1")
}

func main() {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		seq, _ := strconv.Atoi(c.Query("seq"))
		c.String(http.StatusOK, "pong-%d", seq)
	})
	server := endless.NewServer(":9090", router)
	server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1] = append(server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1], preSigUsr1)
	server.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1] = append(server.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1], postSigUsr1)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("listen and serve error:", err.Error())
	}
}
