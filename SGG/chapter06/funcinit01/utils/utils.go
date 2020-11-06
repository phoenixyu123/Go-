package utils

import "fmt"

var Age int
var Name string

//Age Name 在用之前(main.go)想初始化这两个变量
func init() {
	Age = 100
	fmt.Println("<utils init>") //比main包内所有函数都先执行
	Name = "Tom"
}
