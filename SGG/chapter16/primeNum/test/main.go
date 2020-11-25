package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	start := time.Now().Unix()

	func() {
		for i := 3; i < 800000; i++ {
			flag := true
			for j := 2; j < int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					flag = false
					break
				}
			}
			if flag == true {
				fmt.Printf("%v ", i)
			}
		}
	}()
	end := time.Now().Unix()
	fmt.Println(end - start)

}
