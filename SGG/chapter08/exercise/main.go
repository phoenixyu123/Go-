package main

import "fmt"

func main() {

	var score [3][5]float64

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("input score[%v][%v]\n", i, j)
			fmt.Scanln(&score[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		sum := 0.0
		ary := 0.0
		for j := 0; j < 5; j++ {
			sum += score[i][j]
		}
		ary = sum / float64(len(score))
		fmt.Printf("第%v个班总分:%v   平均分:%.2f\n", i, sum, ary)
	}

	fmt.Println(score)

}
