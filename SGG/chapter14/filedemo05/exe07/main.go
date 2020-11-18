package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int //统计不同字符个数
}

func main() {
	//打开一个文件创建一个reader
	//每读取一行统计该行有多少个英文数字空格和其他字符

	fileName := "C:/study/GoList/abc.txt"
	file, err := os.Open(fileName)
	//缺点，不能进行读写控制
	//优点，简单
	if err != nil {
		fmt.Printf("\n")
		return
	}

	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	//开始循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //终止符\n
		//读到的字符串给str
		if err == io.EOF { //读到文件末尾退出
			break
		}
		for _, v := range str {
			fmt.Print(string(v))
			switch { //此处不用字符匹配当作if
			case v >= 'a' && v <= 'z':
				count.ChCount++
			case v >= 'A' && v <= 'Z':
				count.ChCount++

			case v >= '0' && v <= '9':
				count.NumCount++

			case v == ' ' || v == '\t':
				count.SpaceCount++

			default:
				count.OtherCount++

			}
		}
	}
	fmt.Printf("字符个数：%v	数字个数：%v	空格个数：%v	其他字符个数：%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
