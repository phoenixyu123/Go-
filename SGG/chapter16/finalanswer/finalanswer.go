package finalanswer

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

var writeChanNum int = 10 //写协程开启个数
var writeCount int = 100  //随即个数
var sortChanNum int = 10
var sortCount int = 100

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeDataToFile(filePathTitle, filePathSuf string, i int, writeChanOk chan<- bool) {
	file, err := os.OpenFile(filePathTitle+strconv.Itoa(i)+filePathSuf, os.O_CREATE|os.O_RDWR, 0666)
	CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < writeCount; i++ {
		writer.WriteString(string(rand.Int63n('z'-'A') + 'A'))
	}
	writer.Flush()
	writeChanOk <- true
}

func sortDataToChan(filePathTitle, filePathSuf string, i int, dataChan chan string, sortChanOk chan<- bool) {
	file, err := os.OpenFile(filePathTitle+strconv.Itoa(i)+filePathSuf, os.O_CREATE|os.O_RDWR, 0666)
	CheckError(err)
	defer file.Close()
	//使用int切片进行排序***************************************
	intSlice := &[]int{} //注意此处定义，看看以往的intSlice:=[]int{1,2,3,4}
	reader := bufio.NewReader(file)
	for j := 0; j < writeCount; j++ {
		data, _ := reader.ReadByte()             //一个一个字节读
		*intSlice = append(*intSlice, int(data)) //注意此处调用切片
	}
	sort.Ints(*intSlice)
	str := ""
	for j := 0; j < writeCount; j++ {
		str += string(byte((*intSlice)[j])) //看看此处每一步
	}
	dataChan <- str //放入通道收工
	sortChanOk <- true

}

func writeDataToNewFile(filePathTitle, filePathSuf string, i int, v string, newWriteChanOk chan<- bool) {
	file, err := os.OpenFile(filePathTitle+strconv.Itoa(i)+filePathSuf, os.O_CREATE|os.O_RDWR, 0666)
	CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("\n" + v)
	fmt.Println(v)
	writer.Flush() //这一句老忘
	newWriteChanOk <- true

}

func main() {

	rand.Seed(time.Now().UnixNano())
	//文件写入状态
	writeChanOk := make(chan<- bool, writeChanNum) //这里引入新概念，chan<-bool表示只能写入bool
	//而不可以进行<-writeChanOk操作
	filePathTitle := "C:/study/GoList/src/SGG/chapter16/finalanswer/writeDataFile_" //以writeDataFile开头
	filePathSuf := ".log"                                                           //后缀
	for i := 0; i < writeChanNum; i++ {                                             //开启十个协程
		go writeDataToFile(filePathTitle, filePathSuf, i, writeChanOk)
	}
	for {
		if len(writeChanOk) == writeChanNum {
			close(writeChanOk)
			break
		}
	}
	//数据通道
	dataChan := make(chan string, 10)
	//协程状态
	sortChanOk := make(chan<- bool, 10)
	//2.当writeDataFile完成写1000个数据到文件后，让sort协程从文件中读取1000个数据并排序，
	//3.考察点：协程和管道+文件的综合使用
	for i := 0; i < sortChanNum; i++ {
		go sortDataToChan(filePathTitle, filePathSuf, i, dataChan, sortChanOk)
	}
	for {
		if len(sortChanOk) == 10 {
			close(dataChan)
			close(sortChanOk)
			break
		}
	}
	//将结果写入新文件
	newWriteChanOk := make(chan<- bool, 10)
	//4.功能拓展开10个协程writeDataTofile，每隔协程随机生成1000个数据，存放到10个文件
	//5.当10个文件都生成了，让10个sort协程从10文件中读取1000个数据，并完成排序重新写入到10个结果文件。
	var i int = 0
	for v := range dataChan {

		go writeDataToNewFile(filePathTitle, filePathSuf, i, v, newWriteChanOk)
		i++
	}
	for {
		if len(newWriteChanOk) == 10 {
			close(newWriteChanOk)
			break
		}
	}

}
