package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binaryFind(a *[10]int, left int, right int, findx int) {
	if left == right && a[left] != findx {
		fmt.Println("NOT FIND!")
		return
	}
	var middle int = (left + right) / 2

	if a[middle] == findx {
		fmt.Printf("FIND IT:a[%v] == %v\n", middle, a[middle])
	} else if a[middle] < findx {
		binaryFind(a, left, middle-1, findx)
	} else {
		binaryFind(a, middle+1, right, findx)
	}
}

func insertIt(a *[10]int, insert int) [11]int {
	var b [11]int
	var index int = -1
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] <= insert {
			b[i+1] = a[i]
		} else {
			b[i+1] = insert
			index = i
			break
		}
	}
	fmt.Println("index = ", index)
	for i := 0; i <= index; i++ {
		b[i] = a[i]
	}
	if index == -1 {
		b[0] = insert
	}
	return b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(101)
	}
	fmt.Println("随机数组:", a)
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] < a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	var x int = 0
	fmt.Println("排序数组:", a)
	fmt.Println("查找:")
	fmt.Scanln(&x)
	binaryFind(&a, 0, len(a)-1, x)

	var insert int = 0
	fmt.Println("插入:")
	fmt.Scanln(&insert)
	b := insertIt(&a, insert)
	fmt.Println("插入前:", a)
	fmt.Println("插入后:", b)
}
