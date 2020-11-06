package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	a3 := make([]int, 3, 3) //长度、容量
	copy(a3, a1)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//从切片中删除元素
	a4 := []int{1, 2, 3, 4, 5, 6, 7}
	//要删除第三个元素3
	a4 = append(a4[:2], a4[3:]...) //使用...拆开
	fmt.Println(a4)
}
