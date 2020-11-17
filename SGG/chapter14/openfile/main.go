package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//func OpenFile(name string , flag int,perm FileMode)(file *File,err error)
//name：文件名 flag打开模式**** perm
//文件打开模式：flag
// const (
//     O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
//     O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
//     O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
//     O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
//     O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
//     O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
//     O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
//     O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
// )
//可以以多个模式打开
//FileMode//指定控制文件的权限，在Linux和Unix下使用
//r->4,w->2,x->1//Linux学过

func main() {

	filePath := "C:/study/GoList/222.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//写入5句
	str := "hello,Gardon\n"

	writer := bufio.NewWriter(file) //带缓存的*Writer

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer带缓存因此还没写到磁盘里，WriteString内容实际是先暂时写入缓存的
	//写入磁盘语句如下Flush
	writer.Flush() //没有这句缓存内的内容就不会被写到磁盘中去

	//注意重复运行会不断从第一行开始写入，进行覆盖

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
	//特别注意最开始OpenFile函数中传入的打开模式flag！！！！，没有读的权限
	//因此此段代码读不到东西,将打开模式从WRONLY改成RDWR就可以正常运行
}
