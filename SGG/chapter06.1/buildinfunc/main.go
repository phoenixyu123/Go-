package main

import "fmt"

func main() {
	n2 := new(int)
	*n2 = 100
	fmt.Printf("%T       %v      %v        %v\n", n2, &n2, *n2, n2)

}
