package main

import "fmt"

func main() {
	//'/'&'%'
	fmt.Println(10 / 4) //2，都是整数，向下取整
	var n1 float32
	var n2 float32 = 10 / 4.0
	n1 = 10 / 4
	fmt.Println(n1) //2
	n1 = 10.0 / 4
	fmt.Println(n1, n2) //2.5

	// %算法
	//公式：a % b = a - a/b * b
	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3) // -10 - (-10)/3 * 3,[!]注意：10/3 = 3,向下取整
	fmt.Println("10%-3=", 10%-3) //10 - 10/(-3) * (-3)= 10+3*(-3)
	fmt.Println("-10%-3=", -10%-3)

}
