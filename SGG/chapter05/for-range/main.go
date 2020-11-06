package main

import "fmt"

func main() {
	var a string
	fmt.Scanln(&a)
	for index, value := range a {
		fmt.Printf("index = %d, value = %c\n", index, value)
	}
	//如果字符串有中文，传统遍历方式会乱码：因为一个汉字UTF-8三个字节
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%c", a[i])
	// }
	// fmt.Println()

	//解决方法，把string转成rune（切片）
	a2 := []rune(a)
	for i := 0; i < len(a2); i++ {
		fmt.Printf("%c", a2[i])
	}

}
