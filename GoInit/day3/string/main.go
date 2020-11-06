package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello"
	//字符单引号
	c1 := 'h'
	c2 := '1'
	c3 := '沙' //
	fmt.Println(s, c1, c2, c3)
	//1字节=8bit
	//'A'=1字节
	//1 utf-8('沙')=3字节
	fmt.Printf("''\n")

	////***********路径地址
	path := "C:\\Go\\src" //单个斜线相当于转义'G'&'s',2个'\'告诉程序只输出斜线
	fmt.Printf("%s\n", path)

	//********
	path2 := "'C:\\Go\\src'"
	fmt.Printf("%s\n", path2)

	s2 := `
		是轻薄
		人情恶
		与送黄昏花易洛
	` //////////////////反引号输出多行
	fmt.Printf("%s\n", s2)
	s3 := `C:\study\GoList\src` //反引号直接输出
	fmt.Printf("%s\n", s3)
	fmt.Println(len(s3))

	//字符串拼接
	name := "mqx"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	///分割+++++++++
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//查询包含
	fmt.Println(strings.Contains(ss, "mq"))
	fmt.Println(strings.Contains(ss, "mm"))

	//前缀
	fmt.Println(strings.HasPrefix(ss, "m")) //true
	//后缀
	fmt.Println(strings.HasSuffix(ss, "is")) //false
	//判断元素最初/最后出现的位置（从0开始
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))     //2
	fmt.Println(strings.LastIndex(s4, "b")) //5

	//拼接
	fmt.Println(strings.Join(ret, "+")) //from[C: study GoList src] to[C:+study+GoList+src]
}
