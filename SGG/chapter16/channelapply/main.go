package main

import "fmt"

func writeData(intChan chan int) { //通道是引用类型
	for i := 0; i < 50; i++ {
		//放入数据
		// fmt.Printf("写入%v\n", i)
		intChan <- i
	}
	close(intChan) //关闭写,可以继续读
}

func readData(intChan chan int, boolChan chan<- bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("Read Data:", v)
	}
	//读取完毕
	boolChan <- true
	// close(boolChan) //关闭写
}

func main() {
	var intChan chan int
	intChan = make(chan int, 50)
	var boolChan chan<- bool
	boolChan = make(chan<- bool, 5)

	go writeData(intChan)
	for i := 0; i < 5; i++ {
		go readData(intChan, boolChan)
	}

	//需要设置主线程关闭时间
	//time.Sleep(10*time.Second)

	for {
		if len(boolChan) == 5 {
			close(boolChan)
			break
		}
	}

	// for {
	// 	_, ok := <-boolChan //使用quitChan判断是否读取完毕,没读完就不停循环
	// 	if !ok {
	// 		break
	// 	}
	// }

}
