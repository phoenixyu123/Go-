package main

import "fmt"

func main() {
	//100以内数求和当和第一次大于20的当前数
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println(i)
			break
		}
	}
}
