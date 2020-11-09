package main

import "fmt"

type integer int

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string { //自定义String方法输出字符串
	str := fmt.Sprintf("Name = [%v]  Age = [%v]\n", s.Name, s.Name)
	return str
}

func (i integer) print() {
	fmt.Println("i:= ", i)
}
func (i *integer) print2() {
	(*i)++
	fmt.Println("i:= ", *i)
}

func main() {
	var i integer = 10
	i.print()

	var i2 integer = 20
	i2.print2()

	stu := Student{"Tom", 20}
	fmt.Println(&stu) //必须传地址
	fmt.Println(stu)

}
