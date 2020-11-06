package main

import "fmt"

func exchange(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	a, b = exchange(a, b)
	fmt.Println(a, b)
}
