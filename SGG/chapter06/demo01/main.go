package main

import "fmt"

func main() {
	var n1 float64
	var n2 float64
	var op string
	var res float64
	var operator byte = '+'
	fmt.Println("请输入n1   运算符   n2:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&operator)
	// switch op {
	// case "+":
	// 	res = n1 + n2
	// case "-":
	// 	res = n1 - n2
	// case "*":
	// 	res = n1 * n2
	// case "/":
	// 	res = n1 / n2
	// default:
	// 	fmt.Println("运算符输入错误!")
	// }
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("运算符输入错误!")
	}
	fmt.Println(n1, op, n2, operator)
	fmt.Printf("结果res = %v\n", res)

}
