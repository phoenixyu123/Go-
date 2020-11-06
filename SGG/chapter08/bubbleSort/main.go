package main

import "fmt"

func bubbleSort(a *[5]int) {

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

func main() {
	var a [5]int = [...]int{124, 521, 32, 42, 5}
	bubbleSort(&a)
	for i, v := range a {
		fmt.Printf("a[%v] = %v\n", i, v)
	}
}
