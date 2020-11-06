package main

import "fmt"

func main() {
	//两个数求最大值
	// var n1 int = 10
	// var n2 int = 20
	// var max int = n1
	// if n1 < n2 {
	// 	max = n2
	// }
	// fmt.Println(max)
	//求三个数最大值
	var s1 int = 10
	var s2 int = 20
	var s3 int = 30
	var max int = s1
	if s2 > s1 {
		if s3 > s2 {
			max = s3
		} else {
			max = s2
		}
	} else {
		if s3 > s1 {
			max = s3
		}
	}
	fmt.Println(max)
}
