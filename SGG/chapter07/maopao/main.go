package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Num int = 99

func main() {
	start := time.Now().Unix()
	var a [99]int
	for i := 0; i < Num; i++ {
		a[i] = rand.Intn(Num)
	}

	for i := 0; i < Num-1; i++ {
		for j := 0; j < Num-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp

			}

		}
	}
	fmt.Println(a)
	end := time.Now().Unix()
	fmt.Println("耗时:", end-start)
}
