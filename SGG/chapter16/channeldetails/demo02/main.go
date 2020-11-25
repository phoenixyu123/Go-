package main

import "fmt"

func main() {

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i

	}
	//string通道5个数据
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//如果不关闭会阻塞而导致deadlock
	//实际开发中不好确定什么时候关闭通道
	//可以使用select方式解决！！！！！
	// label:
	for {
		select { //case v:=<-管道
		case v := <-intChan: //注意 管道最终都没有关闭也不会因为阻塞而导致死锁
			//取不到会到下一个channel取，！！！！！！！！！！很关键，灵活用于exitChan
			fmt.Printf("从intChan读取数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从intChan读取数据%v\n", v)
		default:
			fmt.Println("此处加入业务逻辑")
			// break//此处会break select
			// time.Sleep(time.Second)
			return //return直接返回肯定管用
			// break label //不推荐，好处是后面代码还能继续执行
		}
	}
}
