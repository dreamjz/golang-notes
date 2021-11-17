package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", upload)
	router.Run(":9090")
}

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	username := c.Query("username")
	now := time.Now().Format("20060102150405")
	dst := fmt.Sprintf("%s_%s", username, now)
	log.Println("dst:", dst)
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
}
