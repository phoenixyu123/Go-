package main

//此章节学习上下文内容

import (
	"example2/gee"
	"net/http"
)

func main() {
	r := gee.New() //路由映射结构体router=>map[string]Handler

	//注意将Handler参数变成了gee.Context ，提供了查询Query/PostForm参数的功能
	//gee.Context封装了HTML/String/JSON函数，能够快速构造HTTP响应
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		//希望做成 /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s \n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			//PostForm获取http.Request.Form的POST值
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})  

	r.Run(":9999") //start Web

}
