package main

import "fmt"

func test() bool {
	fmt.Println("test...")
	return true
}

func main() {
	var i int = 10
	//短路与&&
	// if i > 9 && test() {
	// 	fmt.Println("ok")
	// }
	// if i < 9 && test() { //不执行test()* 短路与现象
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("false")
	// }

	if i > 9 || test() {
		fmt.Println("ok")//短路或--》》不执行test()
	}
	// var age int = 40
	// if age > 30 && age < 50 {
	// 	fmt.Println("true1")
	// } else {
	// 	fmt.Println("false1")
	// }
	// if age < 30 && age < 40 {
	// 	fmt.Println("true2")
	// } else {
	// 	fmt.Println("false2")
	// }

	// if age > 30 || age < 50 {
	// 	fmt.Println("true3")
	// } else {
	// 	fmt.Println("false3")
	// }
	// if age > 30 || age < 40 {
	// 	fmt.Println("true4")
	// } else {
	// 	fmt.Println("false4")
	// }
}
