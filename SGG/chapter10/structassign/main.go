package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	var stu1 Student = Student{
		"小明",
		20,
	}
	stu2 := Student{"小米", 30}
	//在创建结构体变量时将变量名和变量值写在一起
	//防止以后更改结构体内变量顺序后报错
	var stu3 = Student{
		Name: "Java",
		Age:  202,
	}
	stu4 := Student{
		Age:  10,
		Name: "Mary",
	}
	fmt.Println(stu1, stu2, stu3, stu4)
	var stu5 = &Student{
		Name: "小王",
		Age:  20,
	}
	stu6 := &Student{"AKA", 77}
	var stu7 = &Student{
		"小红",
		22,
	}
	fmt.Println(*stu5, *stu6, *stu7)

}
