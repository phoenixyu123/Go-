package main

import (
	"SGG/chapter16/finalanswer"
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//1.writeToFile十个协程，随机生成1000个数，存放到十个不同的文件里
//2.sortData十个协程，对十个不同的文件进行排序操作，

func writeToFile(fileTitle, fileEnd string, i int, writeChanOk chan<- bool) {
	file, err := os.OpenFile(fileTitle+strconv.Itoa(i)+fileEnd, os.O_CREATE|os.O_WRONLY, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for j := 0; j < 2000; j++ {
		writer.WriteString(strconv.Itoa(rand.Intn(10000)) + "\n")
	}
	writer.Flush() //
	writeChanOk <- true

}

func main() {
	rand.Seed(time.Now().UnixNano())
	fileTitle := "C:/study/GoList/src/SGG/chapter16/finaltest02/Begin_" //文件以Begin_开头
	fileEnd := ".log"                                                   //后缀

	writeChanOk := make(chan<- bool, 10)

	for i := 1; i <= 10; i++ {
		go writeToFile(fileTitle, fileEnd, i, writeChanOk)
	}
	for {
		if len(writeChanOk) == 10 {
			close(writeChanOk)
			break
		}
	}
	
	//排序
	sortChan := make(chan string)
	sortChanOk := 
}
