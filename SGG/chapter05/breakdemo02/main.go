package main

import "fmt"

func main() {
	//指定label使用break
	// label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break label1  //只输出一对
				break //输出四对
			}
			fmt.Printf("j = %v\n", j)
		}
	}
}
