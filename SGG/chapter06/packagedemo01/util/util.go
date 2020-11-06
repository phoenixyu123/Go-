package util

import "fmt"

// Cal is a function
func Cal(n1 float32, n2 float32, op byte) float32 {
	//为了让其他包的文件使用Cal函数，需要C大写，类似其他语言public.
	//[注意]：函数上方的注释是必须的！！！！
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
