package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start time
		t := time.Now()
		// Process request
		c.Next()
		// Calculate process time
		log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
