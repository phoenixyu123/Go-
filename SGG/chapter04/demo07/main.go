package main

import "fmt"

func inputit() (int, string, float32, bool) {
	var age int = 1
	var name string = "xm"
	var salary float32 = 1.1
	var isPass bool = true
	// fmt.Scanln(&age, &name, &salary, &isPass) //全部使用引用传递，把地址传进函数*
	return age, name, salary, isPass
}

type Worker struct {
	age    int
	name   string
	salary float32
	isPass bool
}

func main() {
	var age int
	var name string
	var salary float32
	var isPass bool
	var a Worker = Worker{10, "xm", 10.32, false}
	age, name, salary, isPass = inputit()
	fmt.Println(age, name, salary, isPass)
	fmt.Printf("%+v %+v %+v %+v\n", age, name, salary, isPass)
	fmt.Printf("%+v\n", a)
}
