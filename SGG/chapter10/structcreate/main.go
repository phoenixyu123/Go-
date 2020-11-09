package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//case1
	p1 := Person{"Mary", 20}

	//case2
	var p2 Person

	//case3
	var p3 *Person = new(Person)
	(*p3).Name = "Smith"
	p3.Age = 30

	//case4
	var p4 *Person = &Person{}
	(*p4).Name = "Hippo"
	p4.Age = 99

	fmt.Println(p1, p2, *p3, *p4, p3.Age, (*p3).Age)

}
