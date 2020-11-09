package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c *Circle) area2() float64 {
	// c.radius = 0   //里面修改外面也变化
	fmt.Printf("area2种c的地址:%p\n", c)
	return 3.14 * c.radius * c.radius
}

func main() {

	c := Circle{4.0}

	res := c.area()
	fmt.Println("res=", res)
	c2 := Circle{5.0}
	fmt.Printf("area2种c的地址:%p\n", c2)

	res2 := (&c2).area2()
	res3 := c2.area2()
	fmt.Println("res2:= ", res2)
	fmt.Println("res3:= ", res3)
}
