package main

import (
	"fmt"
	"math"
)

func main() {
	//求ax2+bx+c=0的根
	var a, b, c float64
	fmt.Println("请输入a，b，c")
	fmt.Scanln(&a, &b, &c)
	var x1, x2 float64
	x1 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 = (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	// fmt.Println(b ^ 2) //b = 2 = [0000 0010] ^2=[0000 0010]异或[0000 0010] = [0000 0000]
	if b*b-4*a*c > 0 {
		fmt.Printf("两个不同解:%f , %f", x1, x2)
	} else if b*b-4*a*c == 0 {
		fmt.Printf("两个相同解:%f , %f", x1, x2)
	} else {
		fmt.Println("无实根")
	}
}
