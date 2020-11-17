package main

import (
	"bufio"
	"os"
)

func main() {

	filePath := "C:/study/GoList/222.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	//改为追加到文档末尾
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "附加的文本\r\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
