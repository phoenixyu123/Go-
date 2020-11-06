package main

import "fmt"

func cal(n1 float32, n2 float32, op byte) float32 {
	var res float32
	switch op {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		{
			fmt.Println("[ERROR] TLINK(1109):CANNOT RECOGNIZE THE OPERATOR SIGNAL!")
			return 0
		}
	}
	return res
}

func main() {
	var n1, n2 float32 = 1.2, 2.4
	var op byte = '+'
	res := cal(n1, n2, op)
	if res != 0 {
		fmt.Println(res)
	}

}
