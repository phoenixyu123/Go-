package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface //A接口继承B接口和C接口的方法
	test03()
}

type Student struct {
}

func (s Student) test01() {

}

func (s Student) test02() {

}

func (s Student) test03() {

}

type T interface {
}

func main() {
	var stu1 Student
	var a1 AInterface = stu1 //Student必须实现所有接口的方法
	//少满足一个接口都会报错，不能var AInterface
	a1.test01()
	a1.test02()
	a1.test03()    //注意接口是引用类型，而非值类型，即传递接口变量时就是在传递接口地址
	var t T = stu1 //空接口可以接收任何数据类型
	//认为stu1默认实现了T这样一个空接口的任何方法
	var t2 interface{} = stu1 //等于创建一个t2空接口
	fmt.Println(t, t2)

}
