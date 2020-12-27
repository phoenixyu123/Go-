package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNumFind(intChan chan int, primeChan chan int, exitChan chan<- bool) {

	var num int
	for i := 0; i < 2000; i++ {
		num = <-intChan
		//判断是不是素数
		flag := true
		for i := 2; i < num/2+1; i++ {
			if num == 1 || num == 2 {

				flag = false
				break
			}
			if num%i == 0 {
				flag = false
				// fmt.Printf("找到个%v\n", num)
				break
			}
		}
		if flag == true {
			primeChan <- num
		}
		if len(intChan) == 0 {
			break
		}
	}
	exitChan <- true
	fmt.Println("读协程关闭")
}

func main() {
	intChan := make(chan int, 8000)

	primeChan := make(chan int, 2000)

	exitChan := make(chan<- bool, 4)

	//开启一个协程，向intChan放入1-8000个数

	go putNum(intChan)
	//开四个协程，从intChan取出数据并判断是否为素数
	for i := 0; i < 4; i++ {
		go primeNumFind(intChan, primeChan, exitChan)
	}
	for {
		if len(exitChan) == 4 {
			close(primeChan)
			close(exitChan)
			fmt.Println("读关闭")
			break
		}
	}
	for {
		if len(primeChan) == 0 {
			break
		}
		prime := <-primeChan
		fmt.Printf("%v ", prime)
	}
}
