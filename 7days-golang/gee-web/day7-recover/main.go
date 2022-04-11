package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger(), gee.Recovery())
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"A"}
		c.String(http.StatusOK, "%s", names[3])
	})
	r.Run(":9999")
}
