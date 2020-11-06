package main

import "fmt"

func test01(array [3]int) { //[3]int 和 [4]int 不是一个数据类型 []int 是切片
	array[0] = 88
}

func test02(array *[3]int) {
	array[0] = 888
	(*array)[1] = 8888
}

func main() {
	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println(arr) //arr传入值类型，不会改变内容

	test02(&arr)
	fmt.Println(arr)
}
