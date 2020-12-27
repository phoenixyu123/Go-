package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

//空结构体engine用于实现方法ServeHTTP，参考base1中的http

func (this *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Println("-----------------base2-----------------")
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		fmt.Println("-----------------base2-----------------")
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Println("-----------------base2-----------------")
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	fmt.Println("-----------------base2-----------------")
	engine := new(Engine)

	//迈出了Gee框架的第一步，向ListenAndServe的第二个参数传入实现接口的结构体
	//以此来实现拦截所有http请求，又有了统一的控制入口
	log.Fatal(http.ListenAndServe(":9999", engine)) //去比较base1中的代码
	//log.Fatal(http.ListenAndServe(":9999", nil))
	//由于自己通过engine结构体实现了http的接口ServeHTTP，并
	//修改了http.Request的内容（根据不同的URL.Path），在终端输出不同的内容
}
