package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("-----------------base1-----------------")

	//设置两个路由分别是/和/hello，分别绑定indexHandler和helloHandler
	//根据不同的http请求会调用不同的处理函数
	//访问/，响应是URL.Path = /，而/hello的响应则是请求头header中的键值对的信息。
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
	//用于启用web服务。第一个参数是地址，9999表示端口，第二个参数表示所有的http请求实例
	//nil代表使用标准库中的实例处理。第二个参数则是基于net/http标准库实现Web框架的入口
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header { //返回键值对信息
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
