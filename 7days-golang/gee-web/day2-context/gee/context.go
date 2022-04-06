package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is the alias of map[string]interface{}
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm returns the form value from request
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query returns the query string from request
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status set the http status code for context and response
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader set the http header for response
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// String set response with string format
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text-plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values)))
}

// JSON set response with JSON format
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data write data to response
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML set response with HTML format
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
