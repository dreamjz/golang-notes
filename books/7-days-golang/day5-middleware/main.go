package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {
	r := gee.New()

	r.Use(gee.Logger())

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Welcome v1</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		m1 := func(c *gee.Context) {
			log.Println("Before v2 middleware#1")
			c.Next()
			log.Println("After v2 middleware#1")
		}
		m2 := func(c *gee.Context) {
			log.Println("Before v2 middleware#2")
			c.Next()
			log.Println("After v2 middleware#2")
		}

		v2.Use(m1, m2)
	}

	r.Run(":8000")
}
