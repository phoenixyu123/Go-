package main

import (
	"example/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New() //使用New创建gee实例
	//使用GET方法添加路由“/”“/hello”
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			// Header["User-Agent"] = ["curl/7.55.1"]
			// Header["Accept"] = ["*/*"]
		}
	})
	r.GET("/quattro", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			// Header["User-Agent"] = ["curl/7.55.1"]
			// Header["Accept"] = ["*/*"]
		}
	})
	//最后用Run函数启动web服务器，这里的卤藕只是静态路由，不支持/hello/：name这种动态路由
	r.Run(":9999")
}
