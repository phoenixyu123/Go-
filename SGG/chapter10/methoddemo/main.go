package main

import "fmt"

type A struct {
	Num int
}

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name, "针不戳")
}

func (p Person) calcu() {
	var sum int = 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(p.Name, sum)
}
func (p Person) calcu2(n int) {
	var sum int = 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println(p.Name, sum)
}

func (p Person) getSum(a int, b int) int {
	return a + b
}

func (a A) test() {
	a.Num = 100 //值拷贝，不会影响main函数内的值
	fmt.Println("test() Num=", a.Num)
}

func main() {

	a := A{1}
	a.test()
	fmt.Println(a.Num)

	p := Person{"好人"}
	p.speak()
	p.calcu()
	p.calcu2(10)
	fmt.Println("p.getSum = ", p.getSum(10, 20))
	
}
