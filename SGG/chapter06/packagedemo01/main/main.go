package main

import (
	"SGG/chapter06/packagedemo01/util" //从src目录开始往后写路径
	//import写文件夹名，调用时写包名******************************
	"fmt"
)

func main() {
	var n1, n2 float32 = 1.2, 2.4
	var op byte = '+'
	res := util.Cal(n1, n2, op)
	if res != 0 {
		fmt.Printf("%.2f", res)
	}

}
