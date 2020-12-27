package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp := "C:/study/test2.txt"
	// file, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777) //O_TRUNC无论如何全部删除
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777) //O_APPEND附加到最后
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "你好！附加\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
