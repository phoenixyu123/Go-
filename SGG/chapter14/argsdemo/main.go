package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有:", len(os.Args))
	//遍历args切片，读取命令行输入的参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
}
