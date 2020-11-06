package main

import "fmt"

func main() {
	var score int
	fmt.Scanln(&score)
	switch score {
	case 100:
		fmt.Println("A")
	case 99:
		fmt.Println("B")
	default:
		fmt.Println("D")
	}

}
