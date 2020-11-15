package main

import "fmt"

type AInterface interface {
	Say()
}

type stu struct {
	Name string
}

func (stu stu) Say() {
	fmt.Println("stu Say()")
}

func main() {
	var stu stu
	var a AInterface = stu
	a.Say()
	// fmt.Println(AInterface)
}
