package main

import (
	"SGG/model"
	"fmt"
	//for using ptr 's main.go func/var
	//we need pull the pkg in
	//cuz GOPATH:=".......",-->
) //use _ to hold ont the import name when u not use it

//the RULE of identifier
//Combined by a-z,A-Z,0-9,_
//Dont start with number(0-9):[wrong]:0num,42i,etc.
//Golang Case Sensitive:[wrong]:Num≠num
//Dont use space:[wrong]:std user,i love u,etc.
func main() {
	fmt.Println(model.Hero)
	fmt.Println(model.Hero)
	fmt.Println(model.Hero)
	fmt.Println(model.Hero)
	fmt.Println(model.Hero)
	fmt.Println(model.Hero) //?????????????????????????
	var num int = 10
	var Num int = 20
	fmt.Printf("num = %d,Num = %d\n", num, Num)
	//[!]'_' means a null identifier
	var _ int = 20 //Just for placeholder
	//fmt.Println(_)
	//Cant use _ [error],==>_ is not a variable

	//------------test-------------
	//hello\hello12\1hello\h-b\x h\h_4\_ab\int\
	//float32\_\Abc
	//var hello, hello12, h_4, _ab, Abc int = 10
	//[!]var int is ok,var float32 is ok[!]

	//using ptr 's main.go 's var????????????????????????????????????????????????????????????????????????????

	//fmt.Println(model.heroName) 因为heroName首字母小写--》不可被
	//其他包访问(private),若为大写√
}
