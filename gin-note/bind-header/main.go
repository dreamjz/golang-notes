package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	router := gin.Default()
	router.GET("/bindHeader", bindHeader)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func bindHeader(c *gin.Context) {
	var h testHeader
	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, h)
}
