package main

import (
	"fmt"
	"math"
	"time"
)

//求素数

func main() {
	start := time.Now().Unix()
	var n int = 100000
	var flag int = 0
	var sum int = 0
	for i := 2; i < n; i++ {
		s := math.Sqrt(float64(i))
		for j := 2; j < int(s); j++ { //j<i->17s   j<i/2+1->11s  j<sqrt(i)->3s
			if i%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			fmt.Printf("%v ", i)
			sum++
		}
		if sum == 5 {
			sum = 0
			fmt.Println()
		}
		flag = 0
	}
	end := time.Now().Unix()
	fmt.Printf("耗时 %v 秒", end-start)
}
