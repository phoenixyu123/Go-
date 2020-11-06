package main

import "fmt"

func test(n1 *int, n2 int) {
	*n1 = *n1 + 1
	fmt.Printf("*n1的地址：%v , n1的地址: %v\n", n1, &n2)
	fmt.Println("n1=", *n1)
}

func getSum(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	return sum, 10
}

func main() {
	//test
	n1 := 10
	test(&n1, n1)
	fmt.Println("n1(main)=", n1)
	// fmt.Println("*n1(main)的地址：%v\n", &n1)
	n2 := 20
	sum, n1 := getSum(n1, n2)
	fmt.Println("main sum = ", sum)
	fmt.Println("main n1 = ", n1)
}
