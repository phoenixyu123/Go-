package main

import "fmt"

func main() {
	str := "hello@atguigu"

	slice := str[6:]
	fmt.Println("slice =", slice)

	// str[0] = 'z'  //不可此方式改变字符串
	slice2 := []byte(str)
	slice2[0] = 'z'
	str = string(slice2)
	fmt.Println("str =", str)
	//不能此方式处理汉字

	str2 := "针不戳针不戳"
	slice3 := []byte(str2)
	slice3[0] = 'z'
	str2 = string(slice3)
	fmt.Println("str2 = ", str2) //乱码
	str3 := "针不戳针不戳"
	slice4 := []rune(str3) //rune是按字符进行处理，兼容汉字
	slice4[0] = 'z'
	str3 = string(slice4)
	fmt.Println("str3 = ", str3)
}
