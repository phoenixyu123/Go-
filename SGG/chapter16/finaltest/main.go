package main

import (
	"SGG/chapter16/finalanswer"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var goRoutineNum int = 1
var randCount int = 20
var randNum int = 20

func writeDataToFile(filePath string, numChan chan int, writeChanOk chan<- bool) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i := 0; i < randCount; i++ {
		num := rand.Intn(randNum)
		writer.WriteString(strconv.Itoa(num) + "\n")
		numChan <- num
	}
	writer.Flush()
	writeChanOk <- true
}

func getDataFromChan(filePath string, numChan chan int, sortChanOk chan<- bool, sortSlice [20]int) {
	var i int = 0
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for {
		// str := <-numChan
		// num, err := strconv.Atoi(str)
		// finalanswer.CheckError(err)
		if len(numChan) == 0 {
			break
		}
		sortSlice[i] = <-numChan
		fmt.Println(sortSlice[i])
		writer.WriteString(strconv.Itoa(sortSlice[i]) + "\n")
		i++

	}
	writer.Flush()
	sortChanOk <- true

}

func main() {

	rand.Seed(time.Now().UnixNano())
	writeChanOk := make(chan<- bool, goRoutineNum)
	numChan := make(chan int, randCount)
	fptest := "C:/study/GoList/src/SGG/chapter16/finaltest/test.log"
	fpsort := "C:/study/GoList/src/SGG/chapter16/finaltest/sort.log"

	go writeDataToFile(fptest, numChan, writeChanOk)

	for {
		if len(writeChanOk) == 1 {
			close(writeChanOk)
			close(numChan) //不可写入只可读
			break
		}
	}
	sortChanOk := make(chan<- bool, 2)
	var sortSlice [2][20]int
	for i := 0; i < 2; i++ {
		go getDataFromChan(fpsort, numChan, sortChanOk, sortSlice[i])
	}
	for {
		if len(sortChanOk) == 2 {
			close(sortChanOk)
			break
		}
	}
	// for i := 0; i < randCount; i++ {
	// 	reader:=os.OpenFile(fpsort,)
	// }
	// var intSlice []int
	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 20; j++ {
	// 		intSlice = append(intSlice, sortSlice[i][j])
	// 	}
	// }
	// sort.Ints(intSlice)
	// for _, v := range intSlice {
	// 	fmt.Println(v)
	// }
}
