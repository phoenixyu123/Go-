package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	close(intChan) //关闭通道,不能再写入数据
	// intChan <- 3//panic():send on closed channel
	n1 := <-intChan
	fmt.Println(n1)
	fmt.Println("ok")

}
