package main

import "fmt"

func main() {
	var hen []int
	fmt.Println("请输入")
	var n int = 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&hen[i])
	}

}
