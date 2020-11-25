package main

import (
	"fmt"
	"math"
	"time"
)

var goRoutineNum int = 8
var primeArea int = 800000

func putDataToChan(intChan chan int) {
	for i := 3; i <= primeArea; i++ {
		intChan <- i
	}
	close(intChan)
}

func findPrimeNum(intChan chan int, i int, primeChan chan int, exitChan chan<- bool) {
	var flag bool
	for {
		num, ok := <-intChan
		flag = true
		if !ok {
			break
		}
		for j := 2; j < int(math.Sqrt(float64(num))); j++ {
			if num%j == 0 {
				flag = false
				break
			}
		}
		if flag == true {
			primeChan <- num
		}
	}
	// fmt.Printf("第%v个协程取不到数据退出了！\n", i)
	exitChan <- true

}

func main() {

	intChan := make(chan int, primeArea/2)
	primeChan := make(chan int, primeArea/2)
	exitChan := make(chan<- bool, goRoutineNum)
	start := time.Now().Unix()
	go putDataToChan(intChan)

	for i := 1; i <= goRoutineNum; i++ {
		go findPrimeNum(intChan, i, primeChan, exitChan)
	}
	for {
		if len(exitChan) == goRoutineNum {
			close(exitChan)
			close(primeChan)

			break
		}
	}
	for {
		if num, ok := <-primeChan; !ok {
			break
		} else {
			fmt.Printf("%v ", num)
		}
	}
	end := time.Now().Unix()
	fmt.Println(end - start) //耗时

}
