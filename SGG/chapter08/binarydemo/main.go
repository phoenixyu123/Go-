package main

import "fmt"

func binaryFind(arr *[16]int, left int, right int, findx int) {

	mid := int((left + right) / 2)
	if (left == right || left > right) && arr[left] != findx {
		fmt.Println("NOT FOUND!")
		return
	}
	if arr[mid] == findx {
		fmt.Printf("FIND IT: arr[%v] = %v\n", mid, arr[mid])
	} else {
		if arr[mid] > findx {
			binaryFind(arr, left, mid-1, findx)
		} else {
			binaryFind(arr, mid+1, right, findx)
		}
	}
}

func main() {
	arr := [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 23, 42, 52, 64}

	var findx int = 0
	fmt.Println("input")
	fmt.Scanln(&findx)
	binaryFind(&arr, 0, len(arr)-1, findx)
}
