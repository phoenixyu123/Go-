package main

import (
	"fmt"
)	

func main() {
	// desTest()
	
	desTest1()
}
func desTest1() {
	fmt.Println("==========des加解密==========")
	fmt.Println("请输入:")
	var src []byte
	fmt.Scanln(&src)
	// src = paddingText1(src, 8)
	// fmt.Println(src)
	// unpaddingText1(src)
	key := []byte("12345678")
	// encryptDES1(src, key)
	str := encryptDES1(src, key)
	fmt.Println(string(str))
	str = decrypteDES1(str, key)
	fmt.Println("解密数据:" + string(str))

}

func desTest() {
	fmt.Println("==========des加解密==========")
	fmt.Println("请输入:")
	var src []byte
	fmt.Scanln(&src)

	//src := []byte("少壮不努力，老大徒伤悲.")
	key := []byte("12345678")
	str := encryptDES(src, key)
	fmt.Println(str)
	str = decrypteDES(str, key)
	fmt.Println("这是解密数据：" + string(str))
}
