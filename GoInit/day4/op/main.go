package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2
	)
	a++
	if a == b {
		fmt.Println(a, b)

	} else {
		fmt.Println("a!=b")
	}

	//位运算：2进制
	//&与运算
	fmt.Println(5 & 2) //101/10
	//|按位或
	fmt.Println(5 | 2)
	//^异或
	fmt.Println(5 ^ 2)
	//<<左移
	fmt.Println(5 << 1) //101->1010=10

}
