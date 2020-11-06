package main

import (
	"SGG/chapter06/funcinit01/utils"
	"fmt"
)

var age = test()

//为了看到全局变量是先被初始化的,写一个函数
func test() int { //test在init之前执行!,说明先执行全局变量-->init函数-->main函数
	fmt.Println("<test>")
	return 90
}

//init函数在main函数之前就执行完毕,常用于初始化工作
func init() {
	fmt.Println("<init>")
}

func main() {
	fmt.Println("<main>")
	fmt.Println("Age:=", utils.Age) //在utils包中已经被初始化了
	fmt.Println("Name:=", utils.Name)
}
