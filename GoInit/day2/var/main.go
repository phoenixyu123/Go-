package main

import "fmt"

//单独声明
//声明的变量必须使用
var name string
var age int
var isOK bool

// var {
// 	name string
// 	age int
// 	isOK bool
// }批量声明

//声明格式推荐:var studentName string

//***********匿名变量（如函数返回多个值又不想要的
func foo() (int,string){
	return 19, "hanbing"
}

func main() {
	name = "雄是"
	age = 16
	isOK = true                 //变量的赋值
	fmt.Printf("name:%s\n", name) //Println带换行符,或\n
	fmt.Println(name)
	fmt.Println(age)
	//GO语言声明的变量必须使用
	//声明变量同时赋值
	var s2 string = "wanghai"
	fmt.Println(s2)

	//类型推导
	var s3 = "skf"
	fmt.Println(s3)


	//************简短变量声明（只能在函数内使用）
	s4:="haha"//只能在函数内用
	fmt.Println(s4)
	x,_:=foo()////////////////匿名变量
	_,y:=foo()
	fmt.Println(x,y)
}

//调用go fmt main.go 调整格式
