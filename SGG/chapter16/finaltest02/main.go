package main

import (
	"SGG/chapter16/finalanswer"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

//1.writeToFile十个协程，随机生成1000个数，存放到十个不同的文件里
//2.sortData十个协程，对十个不同的文件进行排序操作，
var randCount int = 10
var goroutineCount int = 10

func writeToFile(fileTitle, fileEnd string, i int, writeChanOk chan<- bool) {
	file, err := os.OpenFile(fileTitle+strconv.Itoa(i)+fileEnd, os.O_CREATE|os.O_WRONLY, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for j := 0; j < randCount; j++ {
		writer.WriteString(strconv.Itoa(rand.Intn(10000)) + "\n")
	}
	writer.Flush() //
	writeChanOk <- true

}

func ByteToInt(data []byte) int { ////////////////////字节转int
	bytesBuffer := bytes.NewBuffer(data)
	var x int32 = 0
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}
func Pow(x, y int) int {
	for j := 1; j < y; j++ {
		x = x * 10
		// fmt.Printf("第%v次次方，结果：%v")
	}
	return x
}

func fromByteToInt(data []byte) int {
	end := 0
	for j := 0; j < len(data); j++ {
		// fmt.Println(int(data[j]) - 48)
		end = end + Pow(int(data[j])-48, len(data)-j)

	}
	// fmt.Println(end)//此处经验证输出正确
	return end
}

func sortFromFile(fileTitle, fileEnd string, i int, dataChan chan []int, sortChanOk chan<- bool) {
	file, err := os.OpenFile(fileTitle+strconv.Itoa(i)+fileEnd, os.O_CREATE|os.O_RDONLY, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	intSlice := &[]int{} //
	for j := 0; j < randCount; j++ {
		data, _, _ := reader.ReadLine()
		// fmt.Println(data)[55 55 51 48]...
		intdata := fromByteToInt(data)
		// fmt.Println(end)
		*intSlice = append(*intSlice, intdata)

	}
	sort.Ints(*intSlice)   //排序从小到大
	fmt.Println(*intSlice) //经验证准确无误
	// for _, v := range *intSlice {//不能这么传入通道中！！！！！！！具体原因未知
	// 	dataChan <- v
	// }
	// for j := 0; j < len(*intSlice); j++ {
	// 	dataChan <- (*intSlice)[j]
	// }
	dataChan <- *intSlice //无法实现循环放入通道，索性直接把数组放进去了

	sortChanOk <- true
	// fmt.Println("写入")
}

func endWriteToFile(newFileTitle, fileEnd string, i int, v []int,
	endChanOk chan<- bool) {
	file, err := os.OpenFile(newFileTitle+strconv.Itoa(i)+fileEnd,
		os.O_CREATE|os.O_WRONLY, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	intSlice := v
	writer := bufio.NewWriter(file)
	for j := 0; j < len(intSlice); j++ {
		str := strconv.Itoa(intSlice[j])
		writer.WriteString(str + "\n")
	}
	writer.Flush() //
	endChanOk <- true

}

func main() {
	rand.Seed(time.Now().UnixNano())
	fileTitle := "C:/study/GoList/src/SGG/chapter16/finaltest02/Begin/Begin_" //文件以Begin_开头
	newFileTitle := "C:/study/GoList/src/SGG/chapter16/finaltest02/End/End_"  //文件以End_开头
	fileEnd := ".log"                                                         //后缀

	writeChanOk := make(chan<- bool, goroutineCount)

	for i := 1; i <= goroutineCount; i++ {
		go writeToFile(fileTitle, fileEnd, i, writeChanOk)
	}
	for {
		if len(writeChanOk) == goroutineCount {
			close(writeChanOk)
			fmt.Println("写关闭")
			break
		}
	}

	//排序
	dataChan := make(chan []int, randCount)
	sortChanOk := make(chan<- bool, goroutineCount)
	for i := 1; i <= goroutineCount; i++ {
		go sortFromFile(fileTitle, fileEnd, i, dataChan, sortChanOk)
	}
	for {
		if len(sortChanOk) == goroutineCount { //此处只能得到一个sortChanOk
			fmt.Println("关闭")
			// for i := 0; i < goroutineCount; i++ {
			// 	close(dataChan[i])
			// }
			close(dataChan)
			close(sortChanOk)
			break
		}
	}
	//写入************
	endChanOk := make(chan<- bool, goroutineCount)
	i := 1
	for v := range dataChan {
		go endWriteToFile(newFileTitle, fileEnd, i, v, endChanOk)
		i++
	}
	for {
		if len(endChanOk) == goroutineCount {
			close(endChanOk)
			break
		}
	}
}
