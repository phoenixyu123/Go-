package main

//实现图片拷贝功能
//主要是实现一个能接收源地址和目的地址的函数
//使用带缓存的reader和writer来完成内存的拷贝
//同理可得拷贝视频、音频文件

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写函数接收两个文件路径srcName源文件名字，dstFileName目标文件目录

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcfile, err := os.Open(srcFileName)
	//通过os.Open直接打开源文件，
	if err != nil {
		fmt.Printf("open file err=%v\n", err)

	}
	//通过srcfile ,获取到Reader
	reader := bufio.NewReader(srcfile)
	//按缓存读取srcfile的内容
	dstfile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	//使用OpenFile去读、创建文件
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(dstfile)
	defer srcfile.Close()
	defer dstfile.Close()
	//及时关闭打开的文件对象，会造成内存泄漏
	return io.Copy(writer, reader)

}

func main() {

	//将1.png导入到C:/study/kkk.txt
	//调用Copyfile完成文件拷贝
	srcfile := "C:/study/GoList/1.png"
	dstfile := "C:/study/2.png"
	_, err := CopyFile(dstfile, srcfile)
	if err == nil {
		fmt.Println("拷贝完成")
	}
}
