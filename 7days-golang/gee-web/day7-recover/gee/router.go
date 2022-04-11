package gee

import (
	"net/http"
	"strings"
)

// router 结构
type router struct {
	roots    map[string]*node       // HTTP 方法，前缀树根节点映射
	handlers map[string]HandlerFunc // 路由，处理函数映射
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析路由匹配模式，只能有一个通配符
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, val := range vs {
		if val != "" {
			parts = append(parts, val)
			if val[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 添加路由
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		// 前缀树不存在则创建
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 获取路由
func (r *router) getRoute(method, path string) (*node, map[string]string) {
	// 解析当前请求路径 Path
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	// 搜索前缀树，查找匹配模式
	n := root.search(searchParts, 0)

	if n != nil {
		// 解析匹配模式
		// 因为匹配模式和请求路径解析后的长度相同
		// 可以推断出动态参数或通配参数
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
			if part[0] == '*' && len(part) > 1 {
				// 获取通配符之后的内容
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				// 通配符只能有一个
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// 获取指定方法的所有路由模式
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

// 请求处理
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		// 将路由处理函数追加至 c.handlers
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	// 开始执行
	c.Next()
}
