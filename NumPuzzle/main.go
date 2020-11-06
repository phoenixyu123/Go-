package main

import "fmt"

const (
	a1 = "string"
	a2 = iota
	a3
	a4
	_
	a5
)

func main() {
	fmt.Println(a1, a2, a3, a4, a5)
}
