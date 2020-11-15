package main

import "fmt"

type AInterface interface {
	// a int//error golang 接口中不能包含变量
	Say()
}
type BInterface interface {
	Hello()
}

type stu struct {
	Name string
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("monster Hello()")
}

func (m Monster) Say() {
	fmt.Println("monster Say()")
}

func (stu stu) Say() {
	fmt.Println("stu Say()")
}

func main() {
	var stu stu
	var a AInterface = stu
	a.Say()
	// fmt.Println(AInterface)

	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster //一个结构体变量可以同时满足多个接口
	//golang 接口中不能包含任何变量
	a2.Say()
	b2.Hello()
}
