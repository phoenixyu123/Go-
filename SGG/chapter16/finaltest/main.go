package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func writeDataToFile(randChan chan int, writeQuitChan chan bool) {
	rand.Seed(time.Now().UnixNano())
	fp := "C:/study/GoList/src/SGG/chapter16/finaltest/test.log"
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0666) //os.OpenFile,ioutil知识点
	for i := 0; i < 1000; i++ {
		file.WriteString(strconv.Itoa(rand.Intn(1000)))
		if err != nil {
			fmt.Println("生成随机值失败")
			return
		}
		file.WriteString("\n")
	}
	writeQuitChan <- true
	defer file.Close()
	close(writeQuitChan)

}

func sort(sortQuitChan chan bool, arrChan chan int) {
	fp := "C:/study/GoList/src/SGG/chapter16/finaltest/test.log"
	file, err := os.OpenFile(fp, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("111")//true
	reader := bufio.NewReader(file)
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(str)
	// }
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// fmt.Println(str)
		a := strings.Replace(str, "\n", "", -1) //字符串处理
		if a != "" {
			num, err := strconv.Atoi(a)
			if err != nil {
				fmt.Println(err)
				return
			}
			arrChan <- num
		}

		// fmt.Println(num)
	}
	fmt.Println("读完了")
	// fmt.Println(arr)

	sortQuitChan <- true
	close(sortQuitChan)
	defer file.Close()

}

func main() {
	writeQuitChan := make(chan bool, 1)
	randChan := make(chan int)
	arrChan := make(chan int, 2000)
	sortQuitChan := make(chan bool, 1)
	go writeDataToFile(randChan, writeQuitChan)

	for {
		_, ok := <-writeQuitChan
		if !ok {
			break
		}
	}
	fmt.Println("随机数生成完毕")

	go sort(sortQuitChan, arrChan)

	go func() {
		for i := 0; i < 1; i++ {
			<-sortQuitChan
		}
		close(arrChan)
	}()
	for {
		fmt.Println("loading")
		time.Sleep(time.Second)
	}
	// time.Sleep(2 * time.Second)
	// var arrayInt []int
	// var index int = 0
	// for v := range arrChan {
	// 	arrayInt[index] = v
	// 	index++
	// }

}
