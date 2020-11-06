package main

import "fmt"

func main() {
	var score int
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("A")
	} else if score > 79 && score < 100 {
		fmt.Println("B")
	} else if score > 59 && score < 80 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
