package main

import "fmt"

var (
	//Fun1 : 定义全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	//匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20) //在定义的同时就进行调用了.
	fmt.Println("res1:=", res1)

	//将匿名函数func(n1 int , n2 int ) int 付给a
	//则a为函数类型,可通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println("res2:=", a(10, 30))

	//调用全局匿名函数
	fmt.Println("res3:=", Fun1(4, 9))
}
