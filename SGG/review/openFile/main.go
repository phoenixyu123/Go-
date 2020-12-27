package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp := "C:/study/test.txt"
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//写内容
	str := "testtesttesttest\n"
	//写入时，使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ { //注意此处会直接将原来的内容覆盖
		writer.WriteString(str)
	}

	//因为writer是带缓存，因此调用WriterString
	writer.Flush() //必须加入
}
