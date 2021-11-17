package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	username := c.Query("username")
	for _, file := range files {
		log.Println(file.Filename)
		extension := filepath.Ext(file.Filename)
		filename := strings.TrimSuffix(file.Filename, extension)
		now := time.Now().Format("20060102150405")
		dst := fmt.Sprintf("%s_%s_%s%s", now, username, filename, extension)
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
}
