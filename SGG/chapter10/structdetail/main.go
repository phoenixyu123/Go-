package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	fmt.Println(r1)
	fmt.Printf("r1.leftUp.x=%p\n", &r1.leftUp.x)
	fmt.Printf("r1.leftUp.y=%p\n", &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x=%p\n", &r1.rightDown.x)
	fmt.Printf("r1.rightDown.y=%p\n", &r1.rightDown.y)

	//r2 有两个连续的指针
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Println(r2)
	fmt.Printf("r2.leftUp=%p\n", &r2.leftUp)
	fmt.Printf("r2.rightDown=%p\n", &r2.rightDown)

}
