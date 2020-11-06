package main

import "fmt"

func main() {
	var slice []int = make([]int, 20)
	slice[0] = 10
	slice[2] = 13
	slice[5] = 33
	fmt.Println("slice:=", slice)
	slice = append(slice, 999, 888, 777) //说白了就是数组扩容
	fmt.Println("slice+999:=", slice)
	slice = append(slice, slice...) //必须加...表示切片
	//append在底层创建一个数组，对数组进行扩容，然后让slice变量指向这个数组
	//如果slice = append(slice , ......)则原来的slice切片会在执行后被回收
	
	fmt.Println("slice*2:=", slice)
}
