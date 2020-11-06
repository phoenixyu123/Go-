package main

import "fmt"

func main() {
	var arr [5]int = [...]int{1, 20, 30, 40, 50}
	slice := arr[1:5]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]:=%v\n", i, slice[i])
	}

	for i, v := range slice {
		fmt.Printf("slice[%v]:=%v\n", i, v)
	}

}
