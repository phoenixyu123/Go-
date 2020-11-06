package main

import (
	"SGG/DES/des"
	"fmt"
)

func desfunc() {
	var src []byte
	fmt.Scanln(&src)
	var key []byte = []byte("PhoenixY")
	src = des.EnDes(src, key)
	src = des.DeDes(src, key)
	fmt.Println(string(src))

}

func main() {
	desfunc()
}
