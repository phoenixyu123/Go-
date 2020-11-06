package main

import "fmt"

//常量（程序运行期间不会改变
const pi = 3.1415926535
//批量
const (
	statuOK=200
	notFound=404
)
//不写值同上
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota//0
	a2 = iota//1
	a3//2
	a4//3
)

//test
const (
	b1 = iota//0
	b2
	_
	b3
)

//插队
const (
	c1 = iota//0
	c2 = 100
	c3 = iota
)

//每新增一行+1
const (
	d1,d2 = iota+1,iota+2
	d3,d4 = iota+1,iota+2
)

//*************定义数量级
const (
	_=iota
	KB=1<<(10*iota)//二进制数左移十位
	MB
	GB
	TB
	PB
)

func main(){
	fmt.Println(n1,n2,n3)
	fmt.Println(a1,a2,a3,a4)
	fmt.Println(b1,b2,b3)
	fmt.Println("c:",c1,c2,c3)
	fmt.Println("d:",d1,d2,d3,d4)
	fmt.Println("存储空间:",KB,MB,GB,TB,PB)
}