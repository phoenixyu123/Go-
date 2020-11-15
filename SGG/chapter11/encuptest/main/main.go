package main

import (
	"SGG/chapter11/encuptest/account"
	"fmt"
)

func main() {
	a := account.NewAccount("杨毅刚")
	a.SetPassword("123456")
	a.SetBalance(200)
	fmt.Println(*a)
	fmt.Println(a.GetPassword(), a.GetBalance())
}
