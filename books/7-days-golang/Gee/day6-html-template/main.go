package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Static("/assets", "./static")
	r.LoadHTMLGlob("tmpl/*")

	r.GET("/hello/:name", func(c *gee.Context) {
		c.HTML(http.StatusOK, "tst.tmpl", gee.H{
			"name": c.Param("name"),
		})
	})

	r.Run(":8000")
}
