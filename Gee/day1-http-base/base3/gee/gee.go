package gee

import (
	"fmt"
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

//添加路由需要由功能名method（GET、POST）-请求示例名pattern
//还有handlerfunc实例名（/、/hello）
func (this *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern //如GET-/、GET-/hello、POST-/hello
	this.router[key] = handler
}

//GET函数添加GET请求函数定义
func (this *Engine) GET(pattern string, handler HandlerFunc) {
	this.addRoute("GET", pattern, handler)
}

//POST请求函数定义
func (this *Engine) POST(pattern string, handler HandlerFunc) {
	this.addRoute("POST", pattern, handler)
}

//Run定义了启动http服务器的方法
func (this *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, this) //和base1中第二个参数相比
	//需要将定义好的接口的this结构体类型传入,addr是第一杆参数表示端口号
}

//ServeHTTP定义
func (this *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//Method是Request结构体中的参数，表示如GET\POST等操作
	//空值默认GET操作
	key := req.Method + "-" + req.URL.Path
	if handler, ok := this.router[key]; ok { //映射存在
		handler(w, req) //此时的handler是经过Engine结构体映射后的HandlerFunc请求示例
		//这里注意不用写输出flf，在main包里定义了最终输出
	} else {
		//设置返回码
		w.WriteHeader(http.StatusNotFound) //实际上就是404
		fmt.Fprintf(w, "404 NOT FOUND!: %s\n", req.URL)
	}
}
