package main

import "fmt"

func main() {
	//打印1-100  9的倍数  以及和
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}
