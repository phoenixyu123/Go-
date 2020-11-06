package main

import (
	"SGG/chapter06/exercise02/monkey"
	"fmt"
)

func main() {
	var n int = 0
	var day int
	fmt.Println("请输入：第day天，剩下n个桃子")
	fmt.Scanln(&day, &n)
	fmt.Println(monkey.Monkey(n, day))
	fmt.Println(monkey.Monkey2(day))
}
