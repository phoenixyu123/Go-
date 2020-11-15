package main

import "fmt"

type C struct {
	score int
}

type A struct {
	Name string
	age  int
	C
}

func (a *A) SayOk() {
	fmt.Printf("A.SayOk:%v\n", a.Name)
}
func (a *A) hello() {
	fmt.Printf("A.Hellow:%v\n", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Printf("B.SayOk:%v\n", b.Name)
}

func main() {
	b := B{}
	b.A.Name = "Tom"
	b.A.age = 19
	b.SayOk()
	b.hello() //证明结构体可以使用匿名结构体内所有的方法和属性

	//上面写法可以简化
	b.Name = "Jack"
	b.age = 22 //证明可以直接赋值
	b.score = 99
	b.SayOk()
	b.hello()
	fmt.Println("C.GetScore:", b.score)
}
