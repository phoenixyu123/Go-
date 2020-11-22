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

	//遍历管道不能用普通的for循环+len
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	// for i := 0; i < len(intChan2); i++ {
	// 	fmt.Println(<-intChan2) //只输出前五十个因为每次输出都会使len-1
	// }
	close(intChan2)           //关闭管道后会正常退出
	for v := range intChan2 { //管道没有下标,因此他不能使用下标跳着取
		fmt.Println(v) //报错,取完100个以为还可以取,所以需要用到管道关闭Close(intChan2)
	}

}
