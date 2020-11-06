package main

import "fmt"

//求 a,b最大公约数（欧几里得

func gcd(a int, b int) int {
	if b == 0 {
		return a //递归底层
	} else {
		return gcd(b, a%b)
	}
}

var x1, y1, x, y, r int

func extgcd(a int, b int) (int, int, int) {
	if b == 0 { //递归到底，找到上一层的x1y1的值传回给x和y，上一层a的值就是r（商--》最大公因数
		x1 = 1
		y1 = 0
		x = x1
		y = y1
		r = a
		return r, x, y
	} else {
		r, x1, y1 = extgcd(b, a%b)
		x = y1
		y = x1 - a/b*y1
		return r, x, y
	}
}

func main() {
	var a, b int
	fmt.Println("Please input a & b:")
	fmt.Scanln(&a, &b)
	final := gcd(a, b)
	fmt.Println("Answer:", final)
	fmt.Println("Please input a & b for finding ax+by = gcd(a,b)")
	//fmt.Scanln(&a, &b)
	r, x, y := extgcd(a, b)
	fmt.Println("r,x,y:= ", r, x, y)
	fmt.Printf("ax+by = %d*%d+%d*%d =(gcd(%d,%d)==%d)\n", a, x, b, y, a, b, r)
}
