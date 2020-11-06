package main

import "fmt"

var age int = 50

var Name string = "Jack"

func test() {
	//局部变量
	age := 10
	Name := "Tom"
	fmt.Println(age, Name)

}

func main() {
	test()

}
