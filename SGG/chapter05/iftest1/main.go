package main

import "fmt"

func main() {
	//声明2个int32赋值判断和>50？
	// var a1 int32 = 0
	// var a2 int32 = 0
	// fmt.Scanln(&a1, &a2)
	// if a1+a2 > 50 {
	// 	fmt.Println("HW")
	// }
	// //声明2个float64第一个大于10第二个小于20，打印和
	// var f1 float64 = 0
	// var f2 float64 = 0
	// fmt.Scanln(&f1, &f2)
	// if f1 > 10.0 && f2 < 20.0 {
	// 	fmt.Println(f1 + f2)
	// }

	// //2个数的和能否被3、5整除
	// var b1 int32
	// var b2 int32
	// fmt.Scanln(&b1, &b2)
	// if (b1+b2)%3 == 0 && (b1+b2)%5 == 0 {
	// 	fmt.Println("yes")
	// }
	//判断闰年
	var year int
	fmt.Scanln(&year)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("是闰年")
	}
}
