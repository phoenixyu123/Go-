package main

import "fmt"

//学生案例
//编写student结构体  name，gender，age，id ，score，string，string，int，int，float64
//say方法返回string，包含所有字段

type Student struct {
	Name   string
	Gender string
	Age    int
	id     int
	score  float64
}

func (stu *Student) say() string {
	str := fmt.Sprintf("Name = %v Gender = %v Age = %v Id = %v Score = %v\n",
		stu.Name, stu.Gender, stu.Age, stu.id, stu.score)
	return str
}

func main() {
	stu1 := Student{"Tom", "Male", 18, 1001, 99.09}
	fmt.Println(stu1.say())
}
