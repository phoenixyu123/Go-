package gee

import (
	"net/http"
)

//HandlerFunc 定义gee使用的请求处理程序，在base1和2中都有
type HandlerFunc func(http.ResponseWriter, *http.Request)

//同base2中定义的接口方法一样定义结构体Engine
type Engine struct {
	router map[string]HandlerFunc //定义路由映射表router
	//这里主要是通过映射表可以映射不同的Handler处理方法
}

//使用New函数建立gee.Engine
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	} //返回一个新建的Engine结构体
}

func (this *Engine) addRouter(method string, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern //如GET-/、GET-/hello、POST-/hello
}
