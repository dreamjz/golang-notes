package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	gin.DisableConsoleColor()
	f, err := createLogFile("./log/gin.log")
	if err != nil {
		log.Fatal("create log file failed:", err)
	}
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func createLogFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	log.Println("dir is: ", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("create dir failed:", err)
		}
	}
	return os.Create(path)
}
