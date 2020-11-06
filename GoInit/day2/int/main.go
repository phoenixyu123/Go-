package main

import "fmt"

func main(){
	var i1 = 101
	fmt.Printf("%d %o %x\n", i1,i1,i1)
	//8进制
	i2 :=077
	fmt.Printf("%d\n", i2)
	i3 := 0xfff
	fmt.Printf("%d\n", i3)
}