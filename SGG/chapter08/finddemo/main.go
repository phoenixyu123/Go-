package main

import "fmt"

func main() {
	//一个数列：白猫、金毛、紫衫、青衣
	//从键盘输入名称，判断【顺序查找】是否存在
	var str [4]string = [...]string{"白猫", "金毛", "紫衫", "青衣"}
	var input string
	var flag int = 0
	fmt.Println("input:")
	fmt.Scanln(&input)
	for _, v := range str {
		if input == v {
			fmt.Println("有")
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Println("没有")
	}
}
