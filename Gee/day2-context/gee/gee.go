package gee

import (
	"net/http"
)

//定义框架入口
//HandlerFunc定义了gee使用的请求处理程序
type HandlerFunc func(*Context)

//Engine 实现了ServeHTTP的接口
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	//注意此处使用了函数重载
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	//注意此处使用了函数重载
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
	//注意第二个参数是结构体强转结构体类型
	//只需要在结构体中满足ServeHTTP函数即可
}

//实现ServeHTTP函数，
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req) //构造上下文结构体，它比Req多了三个参数更方便
	engine.router.handle(c)
}

// 将router相关的代码独立后，gee.go简单了不少。
// 最重要的还是通过实现了 ServeHTTP 接口，接管了所有的 HTTP 请求。
// 相比第一天的代码，这个方法也有细微的调整，在调用 router.handle 之前，
// 构造了一个 Context 对象。这个对象目前还非常简单，
// 仅仅是包装了原来的两个参数，之后我们会慢慢地给Context插上翅膀。
