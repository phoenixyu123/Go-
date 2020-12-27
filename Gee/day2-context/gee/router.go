package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

//新建路由结构体变量
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

//添加路由到路由映射表
func (this *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - $s", method, pattern)
	key := method + "-" + pattern
	this.handlers[key] = handler
}

//handle使用context中的函数
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path //别忘了在context上下文中实现的结构体中
	//新增的两个变量req.Method和req.URL.Path

	if handler, ok := r.handlers[key]; ok {
		handler(c) //注意handler是HandlerFunc
	} else {
		c.StatusCode = http.StatusNotFound
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
