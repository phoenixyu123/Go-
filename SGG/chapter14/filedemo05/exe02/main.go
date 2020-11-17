package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个已经存在的文件，并将其修改覆盖
	filePath := "C:/study/GoList/222.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer file.Close()
	str := "你好尚硅谷\n" //\r\n,有的编辑器认\r不认\n所以最好两个都加
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
