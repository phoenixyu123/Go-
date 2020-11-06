package main

import "fmt"

func main() {
	//成绩大于60输出合格，否则不合格
	// var score int = 60
	// fmt.Scanf("%d", &score)
	// switch int(score / 60) {
	// case 0:
	// 	fmt.Println("滚")
	// case 1:
	// 	fmt.Println("针不戳")
	// }

	//打印季节  3 ，4 ，5    6， 7， 8，  9， 10 ， 11 ，    12， 1  ， 2
	var month int
	fmt.Scanf("%d", &month)
	switch month {
	case 3, 4, 5:
		fmt.Println("春")
	case 6, 7, 8:
		fmt.Println("夏")
	case 9, 10, 11:
		fmt.Println("秋")
	case 12, 1, 2:
		fmt.Println("冬")
	default:
		fmt.Println("false")
	}
}
