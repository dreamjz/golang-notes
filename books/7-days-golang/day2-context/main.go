package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Path)
	})

	r.GET("/hello", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":8000")
}
