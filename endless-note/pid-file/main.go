package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		seq, _ := strconv.Atoi(c.Query("seq"))
		c.String(http.StatusOK, "pong-%d", seq)
	})
	server := endless.NewServer(":9090", router)
	server.BeforeBegin = func(addr string) {
		pid := syscall.Getpid()
		log.Println("Actual pid :", pid)
		savePIDFile(pid)
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("listen and serve error:", err.Error())
	}
}

func savePIDFile(pid int) {
	file, err := os.OpenFile("pid", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("create pid file failed:", err.Error())
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(pid))
}
