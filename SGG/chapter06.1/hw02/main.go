package main

import "fmt"

var Monthday []int = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var FishMonth []int = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func cal(y int, m int, d int) int {
	var sum int = 0
	//判断当前年天数      1992 1 1
	sum += d                                    //+1
	m--                                         //1992 0 1
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 { //是闰年
		for i := 0; i < m; i++ {
			sum += FishMonth[i] //sum+=0
		}
	} else {
		for i := 0; i < m; i++ {
			sum += Monthday[i]
		}
	}
	y--
	for i := y; i >= 1990; i-- { //y = 1992 > 1990
		if (i%4 == 0 && i%100 != 0) || i%400 == 0 { //是闰年
			sum += 366
			// fmt.Printf("%v年是闰年+366天,sum=%v", y, sum)
		} else {
			sum += 365
		}
	}
	return sum
}

func stdyltsw(sum int) {
	if sum%5 == 1 || sum%5 == 2 || sum%5 == 3 {
		fmt.Println("打渔")
	} else {
		fmt.Println("晒网")
	}
}

func main() {
	var year int
	var month int
	var date int
	for {
		fmt.Println("\n输入年月日")
		fmt.Scanln(&year, &month, &date)

		fmt.Printf("%v年%v月%v日是自1990年1月1日起的第%v天", year, month, date, cal(year, month, date))
		stdyltsw(cal(year, month, date))
	}

}
