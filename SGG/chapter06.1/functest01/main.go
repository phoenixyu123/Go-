package main

import "fmt"

func jinzita(n int) {
	for i := 0; i < n; i++ {
		for k := 0; k < n-i-1; k++ {
			fmt.Printf(" ")
		}
		for j := 0; j <= 2*i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	jinzita(n)
}
