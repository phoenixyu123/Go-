package main

import "fmt"

func main() {
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 2
	arr01[2] = 2
	//arr01[3] = 999  //ERROR:不能动态定义数组
	fmt.Println(arr01)

}
