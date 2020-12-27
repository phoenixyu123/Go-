package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//演示反射类型

	//通过反射获取传入的变量的type，kind（类别）
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType) //查看reflect返回的类型rType = int

	//2.获取到refect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal) //此然和num相同但是和原来的num意义是不同的
	// n1 := 10
	// n2 := 2 + n1 //errore
	n2 := 2 + rVal.Int()
	fmt.Println(n2)
	fmt.Printf("rVal = %v rVal type = %T\n", rVal, rVal)
	//输出rVal = 10,rVal type = reflect.Value

	//将rVal转成interface{}
	iv := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Printf("num2 type := %T", num2) //int
}

type Student struct {
	Name string
	Age  int
}
type Monster struct {
	Name string
	Age  int
}

//对结构体反射
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	//1.先获取到reflect.Type
	fmt.Printf("rType := %v\n", rType) //main.Student
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal := %v\n", rVal) //{Tom 20}
	//在rVal中取不到name和age
	//做反射一定要用断言
	iv := rVal.Interface()
	// stu := Student{}
	// stu = rVal//error!!!!
	//可以使用类型断言switch iv.(type){case Student:...case Monster:..}
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name := %v\n", stu.Name)
	} //断言成功
	switch iv.(type) { //待检测的类型断言
	case Student:
		fmt.Println("Type is Student")
	case Monster:
		fmt.Println("Type is Monster")
	}

}

func main() {

	//编写案例，对基本类型进行反射的基本操作
	// var num int = 10
	// reflectTest01(num)

	//定义student实例
	// stu := Student{
	// 	Name: "Tom",
	// 	Age:  20,
	// }
	stu2 := Monster{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu2)
}
