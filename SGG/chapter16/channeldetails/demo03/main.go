package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", i)
		time.Sleep(time.Second)
	}
}
func test() {
	defer func() {

		//捕获test抛出的panic
		if err := recover(); err != nil { //recover函数可以接收error
			fmt.Println("test崩溃")
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" //error
	//一个协程崩了其他协程也崩溃了
}

func main() {
	go sayHello()
	go test()
	for {
		fmt.Println("1")
		time.Sleep(time.Second)
	}
}
