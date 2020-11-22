package main

//无法解决多个协程的输出排序问题,如果是一个协程就完全没问题

import (
	"fmt"
	"sort"
	// "sync"
)

var goroutineNum int = 1

func pushInt(numChan chan int) {

	for i := 1; i <= 2000; i++ {
		numChan <- i
		fmt.Printf("正在写入第%v个\n", i)
	}
	close(numChan)
}

func getInt(numChan chan int, quitChan chan bool, resChan chan map[int]int) {
	for {
		n, ok := <-numChan
		if !ok {
			break
		}
		res := 0
		for i := 0; i <= n; i++ {
			res += i
		}

		count := map[int]int{
			n: res,
		}
		// fmt.Println(count)

		resChan <- count
	}
	quitChan <- true
	fmt.Println("退出")
	// close(quitChan)
	// close(resChan)

}

func main() {
	numChan := make(chan int, 2000)
	quitChan := make(chan bool, goroutineNum)
	var resChan = make(chan map[int]int, 2000)
	go pushInt(numChan)
	for i := 0; i < goroutineNum; i++ {
		go getInt(numChan, quitChan, resChan)
	}
	///////////////////////
	go func() {
		for i := 0; i < goroutineNum; i++ {
			<-quitChan
		}
		close(resChan)
	}()

	for v := range resChan { //
		var keys []int
		for k, _ := range v {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, val := range keys {
			fmt.Printf("res[%v]=%v\n", val, v[val])
		}

	}
	// for {
	// 	v, ok := <-resChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("resChan:=%v\n", v) //稍微无序
	// }

	fmt.Println("执行完毕")

}
