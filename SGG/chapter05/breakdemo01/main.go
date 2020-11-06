//随机生成
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//为了生成随机数需要种子
	// rand.Seed(time.Now().Unix()) //返回一个1970的1月1日00:00:00到现在的秒数
	// fmt.Println(time.Now().Unix())
	// n := rand.Intn(100) + 1 //伪随机，永远82
	// fmt.Println(n)
	var num int64 = 0
	var n int64 = 1
	var startTime int64 = time.Now().UnixNano() //UnixNano,纳秒级  Unix,秒级
	var endTime int64
	var count int = 0
	for {
		count++
		if n != 99 {
			num++
			rand.Seed(time.Now().UnixNano()*num ^ n)
			n = int64(rand.Intn(100) + 1)
			fmt.Println(num, n)
		} else {
			break
		}

	}
	endTime = time.Now().UnixNano()
	fmt.Println(endTime-startTime, count)
}
