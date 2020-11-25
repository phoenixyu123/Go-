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
}

func main() {

	//编写案例，对基本类型进行反射的基本操作
	var num int = 10
	reflectTest01(num)
}
