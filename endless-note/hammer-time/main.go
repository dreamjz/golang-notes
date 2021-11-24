package main

import (
	"log"
	"net/http"
	"strconv"
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
	endless.DefaultHammerTime = 1 * time.Second
	server := endless.NewServer(":9090", router)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("listen and serve error:", err.Error())
	}
}
