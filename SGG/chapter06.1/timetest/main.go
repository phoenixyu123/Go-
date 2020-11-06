package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str = str + "Hello" + strconv.Itoa(i+1)
	}
}

func main() {

	//执行test03之前获取时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行十万次拼接用时 %v 秒\n", end-start)

}
