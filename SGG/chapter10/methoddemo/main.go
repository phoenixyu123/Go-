package main

import "fmt"

type A struct {
	Num int
}

func (a A) test() {
	a.Num = 100 //值拷贝，不会影响main函数内的值
	fmt.Println("test() Num=", a.Num)
}

func main() {

	a := A{1}
	a.test()
	fmt.Println(a.Num)
}
