package main

import "fmt"

func swap(n1 *int, n2 *int) {
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}

func main() {
	var n1 int = 10
	var n2 int = 2
	swap(&n1, &n2)
	fmt.Printf("n1 = %v , n2 = %v\n", n1, n2)
}
