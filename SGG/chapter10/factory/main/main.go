package main

import (
	"SGG/chapter10/factory/model"
	"fmt"
)

func main() {
	var stu = model.Student{
		"Tom",
		20,
		233.33,
	}
	fmt.Println(stu)
	// var stu2 = model.student{ //小写不能跨包，使用工厂模式引用
	// 	"Mary",
	// 	30,
	// 	333.33,
	// }
	// fmt.Println(stu2)
	stu2 := *model.NewStudent("Mary", 22, 1.01)
	fmt.Println(stu2)

	fmt.Printf("name = %v , score = %v\n", stu2.Name, stu2.GetScore())
}
