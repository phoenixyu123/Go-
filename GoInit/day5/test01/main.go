package main

import "fmt"

func main() {
	var a1 [5]int
	a1 = [...]int{1, 3, 5, 7, 8}
	var sum int
	for i := 0; i < len(a1); i++ {
		sum += a1[i]
		fmt.Println(sum)
	}
}
