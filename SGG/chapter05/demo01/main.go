package main

import "fmt"

func main() {
	var age int = 0
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println(">18")
	}

	if age = 20; age > 18 {
		fmt.Println(">>18")
	}
}
