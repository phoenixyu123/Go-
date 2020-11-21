package main

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res //true:55 else:false
}
func addUpper2(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res //true:55 else:false
}

func main() {
	res := addUpper(10)
	if res == 55 {
		fmt.Printf("结果=%v		正确\n", res, 55)
	} else {
		fmt.Printf("结果=%v		期望:%v\n", res, 55)
	}
	// fmt.Println(res) //55
}
