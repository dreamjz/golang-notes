package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Run(":9090")
}

func myLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
