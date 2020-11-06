package main

import "fmt"

//转置三维矩阵

func zz(a [3][3]int) [3][3]int {
	var b [3][3]int
	for i := 0; i < 3; i++ {
		b[0][i] = a[i][0]
		b[1][i] = a[i][1]
		b[2][i] = a[i][2]
	}
	return b
}

func main() {
	var a [3][3]int
	fmt.Println("请输入：")
	for i := 0; i < 3; i++ {
		fmt.Scanln(&a[i][0], &a[i][1], &a[i][2])
	}
	a = zz(a)
	for i := 0; i < 3; i++ {
		fmt.Println(a[i][0], a[i][1], a[i][2])
	}
}
