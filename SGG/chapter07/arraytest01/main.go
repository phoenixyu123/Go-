package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//创建一个byte类型的26个元素的数组分别放'A'->'Z',使用for循环访问

	var myCharts [26]byte
	for i := 0; i < 26; i++ {
		myCharts[i] = byte('A' + i)
	}

	for i, v := range myCharts {
		fmt.Println(i+1, string(v))
	}

	//求数组最大值给出下标

	var intArr [6]int = [...]int{1, 23, 2, 14, 43, 29}
	var max int = intArr[0]
	var flag int = 0
	for i := 1; i < 6; i++ {
		if intArr[i] > max {
			max = intArr[i]
			flag = i
		}
	}
	fmt.Println("最大值:", max, " 下标:", flag)

	//求数组平均值
	var sum int = 0
	for _, v := range intArr {
		sum += v
	}
	fmt.Printf("平均值为：%.2f\n", float64(sum)/float64(len(intArr)))

	//随机生成五个数并反转打印
	var intArr2 [5]int
	rand.Seed(time.Now().UnixNano())
	for i, _ := range intArr2 {
		intArr2[i] = rand.Intn(100)
	}
	fmt.Println("初始数组：", intArr2)
	for i := 0; i < len(intArr2)/2; i++ {
		temp := intArr2[i]
		intArr2[i] = intArr2[len(intArr2)-i-1]
		intArr2[len(intArr2)-i-1] = temp
	}
	fmt.Println("反转后:", intArr2)
}
