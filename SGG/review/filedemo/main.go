package main

import (
	"fmt"
	"os"
)

func main() {

	//打开文件
	file, err := os.Open("C:/study/test1.txt")
	if err != nil {
		fmt.Println("open file err")
	}

	fmt.Printf("file = %v", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close err")
	}
}
