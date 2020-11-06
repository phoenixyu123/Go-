package main

import "math"
import "fmt"


func main(){
	// math.MaxFloat32 //float 32 最大值
	fmt.Println(math.MaxFloat32)
	f1:= 1.23456
	fmt.Printf("%T\n", f1)//默认float64
	f2:=float32(1.23456)
	fmt.Printf("%T\n", f2)
	// f2=f1
	// f1=f2      float32和64之间不能相互赋值 
	
	//bool默认false，不允许整型转换成bool==》1≠true
	b1:=true
	fmt.Printf("%T\n", b1)
}