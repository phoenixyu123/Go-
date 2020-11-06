package main

import "fmt"

//打印空心金字塔***********

func main() {

	var starRaw int = 0
	fmt.Println("请输入行数:")
	fmt.Scanln(&starRaw)
	for i := 1; i <= starRaw; i++ {
		for k := starRaw - i; k > 0; k-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*i - 1); j++ {
			if i == starRaw {
				fmt.Printf("*")
			} else {
				if j == 1 || j == (2*i-1) {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		if i != starRaw {
			fmt.Println()
		}

	}
}
