package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	var p1 Person
	fmt.Println(p1)

	// p1.slice[0] = 100

	//slice一定先make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	//map一定先make
	p1.map1 = make(map[string]string)
	p1.map1["TEST"] = "OK"
	fmt.Println(p1)

	p2 := p1 //值拷贝
	p2.Age = 1000
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := &p1
	p3.Age = 999
	fmt.Println(p1, "			", &p1)
	fmt.Println(p2)
	fmt.Println(p3, "			", &p3)
}
