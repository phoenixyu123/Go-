package main

import (
	"SGG/chapter06/packagedemo01/util"
	"SGG/destest"
	"fmt"
)

func main() {
	fmt.Println("test")
	var n1, n2 float32 = 0.1, 0.2
	var op byte = '*'
	fmt.Println(util.Cal(n1, n2, op))
	var s1 []byte = []byte("test")
	var key []byte = []byte("12345678")

	src := destest.EncryptDES1(s1, key)
	key = destest.DecrypteDES1(src, key)
	fmt.Println(string(key))
}
