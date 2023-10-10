package gee

import (
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler function
type HandlerFunc func(*Context)

// Engine is the instance of framework
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// Ensure that *Engine implements the interface
var _ http.Handler = (*Engine)(nil)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Run start http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP conforms to http.Handler interface
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
