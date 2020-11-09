package main

import "fmt"

type Box struct {
	Width  float64
	Length float64
	Height float64
}

type Visitor struct {
	Name string
	Age  int
}

func (v *Visitor) panduan() {
	if v.Age >= 18 {
		fmt.Println(">18")
	} else {
		fmt.Println("free")
	}
}

func (box *Box) getVolumn() float64 {
	return box.Width * box.Length * box.Height
}

func main() {
	box := Box{1.04, 2.02, 3.11}
	fmt.Println(box.getVolumn())

	var v Visitor
	for {
		fmt.Println("输入姓名")
		fmt.Scanln(&v.Name)
		fmt.Println("输入年龄")
		fmt.Scanln(&v.Age)
		v.panduan()
	}
}
