package main

import "fmt"

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
	//a = 1
	//&a = 0x
	//*&a = 1
	//b=&a
	//&b = b的地址
	//*&b = &a
	//*b = 1
	//c = &b*************
	//*c = b = &a
	//&c = &c
	//*&c = c = &b
	//**c = *b = *&a = a =1
	//***&*&*&*&c = **c = **&b = *b = *&a = a =1
	//x = *b = a = 1
}
