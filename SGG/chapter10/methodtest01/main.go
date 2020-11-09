package main

import "fmt"

//编写MethodUtils结构体，编写方法不需要参数，打印10*8的矩形在main调用
type MethodUtils struct {
	// Width int
	// Height int
}

func (m MethodUtils) print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (meth MethodUtils) print2(n, m int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) area(len, width float64) float64 {
	return len * width
}

//判断奇偶
func (m *MethodUtils) panduan(num int) {
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

func main() {
	m := MethodUtils{}
	m.print()

	m2 := MethodUtils{}
	m2.print2(10, 8)

	m3 := MethodUtils{}
	res := m3.area(10, 20)
	fmt.Println("res:=", res)

	m4 := MethodUtils{}
	m4.panduan(99)
}
