package main

import "fmt"

//切片是数组的引用-->切片是引用类型
//切片长度是可变的-->动态变化的数组
//var a []int

func main() {

	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	slice := intArr[1:3] //左闭右开 起始下标：1，终止下标：3不包含3==》intArr[1,2]={2,3}
	fmt.Println("intArr:=", intArr)
	fmt.Println("slice:=", slice)
	fmt.Println("slice长度:=", len(slice))
	fmt.Println("slice容量:=", cap(slice))                                     //一般为2倍的元素量
	fmt.Printf("slice[0]的地址:=%p, intArr[1]的地址:=%p\n", &slice[0], &intArr[1]) //两者相同
	//slice从底层来说就是一个struct
	//slice ---->>>>>{
	// 	Addr=&intArr[0]
	// 	len=len(slice)
	// 	capacity=2*len
	// }
	slice[0]++
	fmt.Println("intArr:=", intArr) //改变slice，intArr就会改变

	//使用make创建切片 var 变量名 []type = make([]type, len, [cap])
}
