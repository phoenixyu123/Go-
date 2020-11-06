package main

import (
	"fmt"
)

func main() {
	//求1-100 素数
	var flag int = 0
	for i := 2; i <= 100; i++ {
		for j := 2; j < i/2+1; j++ {
			if i%j == 0 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			fmt.Printf("%d ", i)
		}
		flag = 1
	}

}
