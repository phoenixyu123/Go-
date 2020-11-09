package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900

}

type Stu struct {
	Name string
	Age  int
	Addr string
}

func main() {

	map1 := make(map[int]int, 2) //自动扩容
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1)

	students := make(map[string]Stu)
	stu1 := Stu{"Tom", 18, "北京"}
	stu2 := Stu{"Mary", 20, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("学生的编号是:%v\n", k)
		fmt.Printf("学生的姓名是:%v\n", v.Name)
		fmt.Printf("学生的年龄是:%v\n", v.Age)
		fmt.Printf("学生的地址是:%v\n", v.Addr)
		fmt.Println()
	}
}
