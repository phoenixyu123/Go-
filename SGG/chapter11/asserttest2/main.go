package main

import "fmt"

type Student struct {
	Name string
}

//判断参数类型
func TypeJudge(items ...interface{}) {
	for index, val := range items {
		switch val.(type) { //type是一个关键字，可以断言是哪个类型
		case bool:
			fmt.Printf("第%v的参数类型是bool，值是%v\n", index, val)
		case string:
			fmt.Printf("第%v的参数类型是string，值是%v\n", index, val)
		case float32:
			fmt.Printf("第%v的参数类型是float32，值是%v\n", index, val)
		case float64:
			fmt.Printf("第%v的参数类型是float64，值是%v\n", index, val)
		case int:
			fmt.Printf("第%v的参数类型是int，值是%v\n", index, val)
		case Student:
			fmt.Printf("第%v的参数类型是Student类，值是%v\n", index, val) //可以断言自己定义的类型
		default:
			fmt.Printf("第%v的参数类型不确定，值是%v\n", index, val)
		}
	}
}

func TypeJudge2(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("other")
		}
	}
}

func main() {

	var n1 float32 = 1.1
	var n2 float64 = 2.222
	var n3 int = 3
	var n4 string = "44444"
	var n5 bool = true
	var n6 = Student{
		Name: "小米",
	}
	TypeJudge(n1, n2, n3, n4, n5, n6)
	TypeJudge2(n1, n2, n3, n4, n5, n6)
}
