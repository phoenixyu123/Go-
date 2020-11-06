package main

import "fmt"

func main() {
	//i=5时跳出for
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	//i=5时跳过此次
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
