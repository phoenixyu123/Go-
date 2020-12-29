package gee

import (
	"log"
	"net/http"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)改为
// type HandlerFunc func(*Context)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}

}

//gee GET时添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s ", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler //同base2添加路由映射

}

//访问时查找路由
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 %s\n", c.Path)
	}
}
