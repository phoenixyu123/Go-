package main

import "fmt"

//make--> 切片
//判断切片为空-->len(S)==0 ,而不是S == nil

func main() {
	s1 := make([]int, 5, 10) //框选连续内存段
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s2, len(s2), cap(s2))

	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	//索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
