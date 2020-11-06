package main

import "fmt"

//AddUpper is
func AddUpper() func(int) int {
	var n int = 10 //					从这一行
	var str = "Hello"
	return func(x int) int { //						|
		n = n + x //								|
		str += string(int('a'-1) + x)
		fmt.Println("str:=", str)
		return n //								|
	} //									到这一行形成一个闭包!!!
}

func main() {
	f := AddUpper() //可以看作f接收了一个AddUpper的返回值
	//而它的返回值是一个匿名函数
	//f就等价于 func(int) int{n = n+x return n}这个函数
	//所以f(1)就是10+1
	fmt.Println("res:=", f(1)) //f(1)它返回的是一个函数,但是执行完之后n就变了
	// (可以看作匿名函数和外面的引用n形成一个整体构成闭包)

	//闭包的关键就是分析清楚返回值的函数和引用到的变量构成一个整体,而不会使引用到的外部变量不会被初始化
	fmt.Println("res:=", f(2)) //因为是闭包,导致上一条语句执行后,f中的n就变成了11
	fmt.Println("res:=", f(3))
}
