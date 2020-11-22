package main

import (
	"fmt"
	"sync"
	"time"
)

//使用goroutine计算1-200各个数阶乘
//1.编写函数计算阶乘并放入到map中
//2.启动多个协程统一的访问map放入结果
//因此map应该做成全局的,能让大家都访问

var myMap = make(map[int]int, 10)
var lock sync.Mutex //sync同步 Mutex互斥(结构体,包含lock和Unlock)

func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {

	for i := 1; i <= 20; i++ {
		go test(i)
	}
	//遍历结果
	//go build -race main.go, main.exe显示found 3 data race(s)资源竞争
	time.Sleep(time.Second) //不加这一句不一定会输出多少个
	lock.Lock()             //注意此处需要加锁,因为myMap为全局变量
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()

}
