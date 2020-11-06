package main

import "fmt"

func main() {
	var height int32
	var money float32
	var handsome bool
	fmt.Scanln(&height, &money, &handsome)
	if height >= 180 && money >= 10000000 && handsome == true {
		fmt.Println("嫁")
	} else if height >= 180 || money >= 10000000 || handsome == true {
		fmt.Println("行")
	} else {
		fmt.Println("滚")
	}
}
