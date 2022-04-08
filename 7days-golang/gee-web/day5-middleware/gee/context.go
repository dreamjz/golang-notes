package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is alias of map[string]interface{}
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string // 路由参数
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index int
}

// the constructor of gee.Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index: -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ;c.index < s; c.index++{
		c.handlers[c.index](c)
	}
}

// Param returns the parameter in URL path
func (c *Context) Param(key string) string {
	val, _ := c.Params[key]
	return val
}

// PostForm returns the parameter in form data
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query returns the parameter in query string
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status set the StatusCode of gee.Context
// And write status code to response
func (c *Context) Status(code int) {
	c.StatusCode = code
	// 写入响应 HTTP Status Code
	c.Writer.WriteHeader(code)
}

// SetHeader set response header
func (c *Context) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
