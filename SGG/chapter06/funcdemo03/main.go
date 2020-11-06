package main

import "fmt"

func test(n int) {
	if n-2 > 0 {
		// fmt.Println(n)
		test(n - 2)
	}
	fmt.Println(n) //考虑跟上面的输出的区别
}

func main() {
	//递归
	var n1 int = 32
	test(n1)
}
