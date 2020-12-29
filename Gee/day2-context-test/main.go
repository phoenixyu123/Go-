package main

import (
	"fak/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello GEE<h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello %s , This is a test text: %s", c.Query("name"), c.Path)
		//String(code int,format string,)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
