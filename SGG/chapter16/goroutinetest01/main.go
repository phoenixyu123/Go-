package main

import "fmt"

var resChan = make(chan int, 2000)

func pushInt(numChan chan int) {

	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func getInt(numChan chan int, quitChan chan bool) {
	for {
		n, ok := <-numChan
		if !ok {
			break
		}
		res := 0
		for i := 0; i <= n; i++ {
			res += i
		}
		resChan <- res
	}
	quitChan <- true
	close(quitChan)

}

func main() {
	numChan := make(chan int, 2000)
	quitChan := make(chan bool, 1)
	go pushInt(numChan)
	go getInt(numChan, quitChan)
	for {
		if _, ok := <-quitChan; !ok {
			break
		}

	}
	fmt.Println("执行完毕")

}
