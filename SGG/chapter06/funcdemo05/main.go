package main

import "fmt"

type myType func(int, int) int //函数变量

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//支持调用函数当形参
func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数变量命名
func myFunc2(funvar myType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数返回值命名(常用***************)
func getSumandSub(n1 int, n2 int) (Sum int, Sub int) {
	Sub = n1 - n2
	Sum = n1 + n2
	return //因为命名了可以不用写了
}

//支持可变参数命名
func sum(n1 int, args ...int) (sum int) {//args可以改名
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func main() {

	a := getSum
	fmt.Printf("a的类型是：%T, getSum的类型是：%T\n", a, getSum)
	res := a(10, 20)
	res2 := myFunc(a, 1, 3)
	fmt.Println(res)
	fmt.Println("res2:=", res2)

	type myInt int //给int取别名,但go认为这两个不是同一个类型，详情见25行
	// var num2 int = 20
	var num1 myInt
	// num1 = num2//报错,可通过转换=>num1 = (myInt)num2
	num1 = 40
	fmt.Printf("num1:=%d\n", num1)
	// b := getSum
	res3 := myFunc2(a, 1, 3)
	fmt.Printf("res3:=%d\n", res3)

	a1, b1 := getSumandSub(2, 4)
	fmt.Printf("a=%v,b=%v\n", a1, b1)

	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("sum:=%v\n", sum)
}
