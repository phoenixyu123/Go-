package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //判断为空nil==>true
	fmt.Println(s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"aa", "bb", "ccc"}
	fmt.Printf("%d %d\n", len(s1), cap(s1)) //长度、容量

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]   //数组切割，左闭右开*==>[0,4)
	fmt.Println(s3) //1,3,5,7

}
