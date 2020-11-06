package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	var slice []int = make([]int, 10)
	fmt.Println("slice=", slice)
	copy(slice, a) //在进行拷贝时数据类型都要是切片*
	//slice和 a两个切片相互独立，改变值不会对另一方产生影响
	fmt.Println("slice after = ", slice)

	var slice2 []int = make([]int, 1)
	fmt.Println("slice2 = ", slice2)
	copy(slice2, a)
	fmt.Println("slice2 after =", slice2)
}
