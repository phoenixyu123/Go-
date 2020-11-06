package main

import "fmt"

func main() {
	var a1 [5]int
	a1 = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
