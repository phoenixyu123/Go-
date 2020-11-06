package main

import "fmt"

func maxtest(a, b int) int { //不能使用max作为函数名
	if a > b {
		return a
	}
	return b //代码规范检查

}

func main() {
	var a int
	var b int
	var c int
	var max int = 0
	fmt.Println("请输入abc:")
	fmt.Scanln(&a, &b, &c)
	max = maxtest(a, b)
	max = maxtest(max, c)
	fmt.Println("最大值：", max)
}
