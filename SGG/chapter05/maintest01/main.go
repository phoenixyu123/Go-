package main

import "fmt"

func main() {
	//从键盘输入不确定个数的数，输入为0结束，判断输入的数有几个正负
	var a [100]int
	var zflag int = 0
	var fflag int = 0
	for i := 0; ; i++ {
		fmt.Println("请输入：")
		fmt.Scanln(&a[i])
		if a[i] > 0 {
			zflag++
		} else if a[i] < 0 {
			fflag++
		} else {
			break
		}
	}
	fmt.Printf("正数 %v 个，负数 %v 个\n", zflag, fflag)
	for i, v := range a {
		fmt.Println(i, v)
	}
}
