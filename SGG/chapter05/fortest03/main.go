package main

import "fmt"

func main() {
	var score [3][5]int
	var sum int = 0
	var passScore int = 60
	var passNum int = 0
	for i := 0; i < 3; i++ {
		fmt.Printf("开始输入第 %v 个班\n", i+1)
		for j := 0; j < 5; j++ {
			fmt.Printf("请输入第 %v 个人的成绩\n", j+1)
			fmt.Scanln(&score[i][j])
			if score[i][j] >= passScore {
				passNum++
			}
		}

	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			sum += score[i][j]
		}
	}
	var ave float32 = float32(sum) * 1.0 / 15.0
	fmt.Printf("三个班平均分: %v\n", ave)
	fmt.Printf("三个班及格人数: %v\n", passNum)

}
