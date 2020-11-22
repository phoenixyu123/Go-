package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println("Test:helloworld" + strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() /////协 程 开 启

	for i := 0; i < 5; i++ {
		fmt.Println("Main:helloworld" + strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}
