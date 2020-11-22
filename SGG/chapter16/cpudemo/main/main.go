package main

import (
	"fmt"
	"runtime"
)

func main() {

	num := runtime.NumCPU()
	fmt.Println("CPU数量=", num)
	//自己设置使用多少cpu

	runtime.GOMAXPROCS(num - 1)

}
