package main

import "fmt"

func test() {
	var n1 int = 10 //n1只能在test中使用
	fmt.Println(n1)
}

func test02(n1 int) {
	n1 = n1 + 10
	fmt.Println("test02 n1:=", n1)
}

func main() {
	n1 := 20
	test02(n1)
	fmt.Println("main n1:=", n1)
}
