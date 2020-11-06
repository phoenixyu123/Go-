package main

import "fmt"

//斐波那契     1,1,2,3,5,8,13...
// func fib1(a int, b int, n int) {
// 	var sum int = a + b
// 	if n-1 > 0 {
// 		fmt.Println(sum)
// 		fmt.Println(float32(b) / float32(sum))
// 		fib1(b, sum, n-1)

// 	} else {
// 		fmt.Printf("fib值为 ： %v\n", sum)
// 	}
// }

func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib2(n-1) + fib2(n-2)
}

func main() {
	var n int
	fmt.Scanln(&n)
	// fib1(1, 1, n)
	fmt.Println("main fib : ", fib2(n))
}
