package main

import "fmt"

//斐波那契数组！（切片
func fib(n int) []uint64 {
	var f = make([]uint64, n)
	f[0], f[1] = 1, 1
	for i := 2; i < n; i++ {
		f[i] = f[i-2] + f[i-1]
	}
	return f
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println("fib:=", fib(n))
}
