package monkey

import "fmt"

//一堆桃子，猴子每天吃一半+1，第十天吃发现没有了，问当初有多少个桃子
//Monkey is a func
func Monkey(n int, day int) (int, int) {
	if day <= 1 {
		return n, day
	} else {
		fmt.Printf("n:%v day:%v\n", n, day)
		return Monkey((n+1)*2, day-1)
	}
}
func Monkey2(day int) int {
	if day < 1 || day > 10 {
		return 0
	} else if day == 10 {
		return 1
	} else {
		return (Monkey2(day+1) + 1) * 2
	}
}
