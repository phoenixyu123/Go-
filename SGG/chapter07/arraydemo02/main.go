package main

import "fmt"

func main() {
	var a1 [3]int = [3]int{1, 2, 3}
	fmt.Println("a1:=", a1)
	var a2 = [3]int{1, 2, 3}
	fmt.Println("a2:=", a2)
	var a3 = [...]int{1, 2, 3}
	fmt.Println("a3:=", a3)
	var a4 = [...]int{1: 1, 0: 2, 2: 3} //指定下标
	fmt.Println("a4:=", a4)

	for i, v := range a1 {
		fmt.Println(i, v)
	}
	

}
