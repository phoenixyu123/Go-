package main

import "fmt"

func main() {
	//switch 小写转大写
	var a byte
	fmt.Scanln(&a)
	switch a {
	case 'a':
		a = 'A'
	case 'b':
		a = 'B'
	default:
		a = '?'
	}
	fmt.Println(a)
}
