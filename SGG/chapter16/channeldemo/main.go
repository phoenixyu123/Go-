package main

import "fmt"

//创建一个可以存放3个int类型的管道
var intChain = make(chan int, 3)

func main() {
	fmt.Printf("intChain的值:%v    类型:%T		他本身的地址:%v\n", intChain,
		intChain, &intChain)
	//发现intChain实际上是一个地址,指向真正管道所在的首地址

	//向管道写入数据
	intChain <- 1
	intChain <- 211
	// intChain <- 333

	//查看管道长度和cap(容量)
	fmt.Println("channel len = %v", len(intChain)) //2
	fmt.Println("channel cap = %v", cap(intChain)) //3 容量不会变化,不允许超容量

	//从管道中读取数据
	var num2 int
	num2 = <-intChain
	fmt.Println("num2 = ", num2)                   //1
	fmt.Println("channel len = %v", len(intChain)) //1
	fmt.Println("channel cap = %v", cap(intChain)) //3 容量不会变化,不允许超容量

	num3 := <-intChain
	fmt.Println("num3 = ", num3)                   //211
	fmt.Println("channel len = %v", len(intChain)) //0
	fmt.Println("channel cap = %v", cap(intChain)) //3 容量不会变化,不允许超容量

	// num4 := <-intChain							  //报错,管道空了
	// fmt.Println("num4 = ", num4)                   //211
	// fmt.Println("channel len = %v", len(intChain)) //0
	// fmt.Println("channel cap = %v", cap(intChain)) //3 容量不会变化,不允许超容量

}
