package main

import "fmt"

func main(){
	var n = 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)//2
	fmt.Printf("%d\n", n)//10
	fmt.Printf("%o\n", n)//8
	fmt.Printf("%x\n", n)//16
	var s = "hello 沙河"
	fmt.Printf("%v\n", s)//同%s
	fmt.Printf("%#v\n", s)//#加双引号
	fmt.Printf("%s\n", s)//同%v

}