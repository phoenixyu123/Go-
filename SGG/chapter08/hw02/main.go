package main

import "fmt"

func main() {

	//跳水打分
	var score [8]int
	for i := 0; i < len(score); i++ {
		fmt.Printf("请第%v个评委打分：\n", i+1)
		fmt.Scanln(&score[i])
	}
	fmt.Println("初次评分：", score)
	var sum int = 0
	var max int = score[0]
	var maxindex int = 1
	var min int = score[0]
	var minindex int = 1
	for i := 0; i < len(score); i++ {
		sum += score[i]
		if max < score[i] {
			max = score[i]
			maxindex = i + 1
		}
		if min > score[i] {
			min = score[i]
			minindex = i + 1
		}
	}
	fmt.Printf("去掉第%v个评委的最高分%v，去掉第%v个评委的最低分%v\n", maxindex, max, minindex, min)
	sum -= (max + min)
	arr := 0.0
	arr = float64(sum) / float64(len(score)-2)
	fmt.Printf("选手最终分数:%.2f", arr)
}
