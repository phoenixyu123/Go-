package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("C:/study/test.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') //读到换行就结束一次
		fmt.Print(str)
		if err == io.EOF { //io.EOF表示文件末尾
			break
		}

	}
	fmt.Println("文件读取结束")
}
