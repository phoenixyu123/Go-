package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//file :=文件指针
	//file :=文件对象
	//file :=文件句柄
	file, err := os.Open("C:/study/GoList/111.txt")
	if err != nil {
		fmt.Println(err)
	}

	//输出文件内容
	fmt.Println(file) //此处输出的是指向文件的指针//对象，都可以
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件错误", err)
	}
}
