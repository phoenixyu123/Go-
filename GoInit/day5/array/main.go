package main

import "fmt"

func main() {
	//数组长度是数组类型的一部分
	var a1 [3]bool //true true false
	var a2 [4]bool //true false false true
	fmt.Printf("%T%T", a1, a2)
	//不初始化
	fmt.Println(a1, a2) //默认全是false

	//初始化方式1
	a1 = [3]bool{true, true, false}
	fmt.Println(a1)

	//初始化方式2
	a10 := [...]int{1, 2, 3, 4, 5} //1 2 3 4 5
	fmt.Println(a10)

	//3
	a3 := [5]int{1, 3} //1 3 0 0 0
	fmt.Println(a3)

	//4
	a4 := [5]int{0: 3, 4: 6} //3 0 0 0 6
	fmt.Println(a4)

	//数组的遍历
	citys := [...]string{"北京", "上海", "广州"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	//2.for-range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多维数组
	a5 := [...][2]int{{1, 2}, {2, 3}}
	fmt.Println(a5)

	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{2, 3},
		[2]int{3, 4},
	}
	fmt.Println(a11)
}
