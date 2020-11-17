package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//要求：读写
	filePath := "C:/study/GoList/222.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	//改为追加到文档末尾
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// err = nil
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	writer := bufio.NewWriter(file)
	str := "附加的文本\n"
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

	// fmt.Println("读取完毕")
}
