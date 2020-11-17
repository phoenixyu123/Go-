package main

//ioutil.ReadFile(*file)
//不需要打开和关闭文件但是没有缓冲

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "C:/study/GoList/111.txt"
	content, err := ioutil.ReadFile(file) //此处没有打开文件

	//因此不需要defer去关闭Close()
	if err != nil {
		fmt.Println(err)
	}
	//输出内容到终端
	fmt.Printf("%v", string(content)) //注意此刻content是[]byte形式

}
