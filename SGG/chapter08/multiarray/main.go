package main

import "fmt"

func main() {
	//二维数组
	var arr [4][6]int = [...][6]int{{0, 0, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0}, {0, 2, 0, 3, 0, 0}, {0, 0, 0, 0, 0, 0}}
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}

	for _, v := range arr {
		// fmt.Printf("i=%v v=%v\n", i, v)
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
}
