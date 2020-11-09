package main

import "fmt"

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (cal *Calcuator) Operator(op byte) float64 {
	res := 0.0
	switch op {
	case '+':
		res = cal.Num1 + cal.Num2
	case '-':
		res = cal.Num1 - cal.Num2
	case '*':
		res = cal.Num1 * cal.Num2
	case '/':
		res = cal.Num1 / cal.Num2
	default:
		fmt.Println("ERROR")
	}
	return res
}

func main() {
	cal1 := Calcuator{1.0, 2.0}
	res1 := cal1.Operator('+')
	res2 := cal1.Operator('-')
	res3 := cal1.Operator('*')
	res4 := cal1.Operator('/')
	fmt.Println(res1, res2, res3, res4)
}
