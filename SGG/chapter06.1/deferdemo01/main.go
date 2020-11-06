package main

import "fmt"

func sum(n1, n2 int) int {
	defer fmt.Println(n1) //暂时不执行，压入defer的独立栈中
	defer fmt.Println(n2) //当这个函数执行完毕后再从defer栈中输出
	//defer栈内的值不会改变，就是存入时候的值
	res := n1 + n2
	n2++
	defer fmt.Println(n2)
	fmt.Println("res:=", res)
	return res
}

func main() {
	//defer的目的：在函数执行后释放defer的资源
	var a int
	var b int
	fmt.Println("请输入：")
	fmt.Scanln(&a, &b)
	a = sum(a, b)
	fmt.Println("main:=", a)
}
