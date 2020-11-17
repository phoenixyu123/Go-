package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("C:/study/GoList/111.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}() //要及时关闭file文件对象，否则会有内存泄漏
	//此处可以直接defer file.Close()，也可以用匿名函数+defer

	//创建一个*Reader，是带缓冲的
	//bufio.NewReader()返回一个*Reader(自带4096字节缓冲大小:defaultBufSize = 4096)
	//缓冲的意思是文件流按4096分批次处理，来读取更大的文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //用reader读取文件的内容，读到一个换行
		//就结束一次读取，实现一行一行的读取**********************
		if err == io.EOF { //io.EOF表示文件的末尾
			break //最后一行读取完毕跳出循环
		}
		//输出内容
		// fmt.Println()//如果用pln会把换行也读进去
		fmt.Print(str)

	}
	fmt.Println("文件读取结束")

}
