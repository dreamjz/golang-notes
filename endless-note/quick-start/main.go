package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main(){
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		seq,_ := strconv.Atoi(c.Query("seq"))
		c.String(http.StatusOK,"pong-%d",seq)
	})
	endless.DefaultReadTimeOut=1
	endless.DefaultWriteTimeOut=1

	err := endless.ListenAndServe(":9090",router)
	if err != nil {
		log.Fatal("listen and serve error:",err.Error())
	}
}
