package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main(){
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.String(http.StatusOK,"pong")
	})
	err := endless.ListenAndServe(":9090",router)
	if err != nil {
		log.Fatal("listen and serve error:",err.Error())
	}
	log.Println("Server on 9090 stopped")
}
