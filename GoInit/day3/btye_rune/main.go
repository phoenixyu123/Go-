package main

import "fmt"

func main() {
	s := "Hello沙河" //11字节*5+3+3
	//len()是求byte字节的数量
	n := len(s)
	fmt.Println(n)
	// for _, c := range s {
	// 	fmt.Printf("%c\n", c) //%c字符
	// }

	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换为rune切片*
	s3[0] = s3[0] + 1
	fmt.Println(string(s3)) //把rune切片强制转换字符串
	c1 := "红"
	c2 := '红' //同rune==>int32类型
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
}
