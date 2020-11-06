package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str2 := "thisis白酒"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i]) //乱码
	}
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i]) //切片
	}
	fmt.Println()

	//字符串转整数----------strconv
	n, err := strconv.Atoi("4645") //Ascii to int
	if err != nil {
		panic(err) //转换错误
	}
	fmt.Println(n + 100)

	//整数转字符串
	n1 := strconv.Itoa(123245)

	fmt.Printf("%T    %v\n", n1, n1)

	//字符串转byte
	var byte1 []byte
	var test1 string = "test1"
	byte1 = []byte(test1)
	fmt.Printf("%T    %v\n", byte1, byte1)

	//byte 转str
	var byte2 []byte = []byte{97, 98, 99}
	var str3 string = string(byte2)
	fmt.Printf("%T      %v\n", str3, str3)

	//十进制转2 , 8, 16 输出格式为字符串
	var str string
	str = strconv.FormatInt(123, 2)
	fmt.Println("123的二进制是：", str)
	str = strconv.FormatInt(123, 8)
	fmt.Println("123的八进制是：", str)
	str = strconv.FormatInt(123, 16)
	fmt.Println("123的16进制是：", str)

	//查找字符串是否在字符串里 strings.Contains
	b := strings.Contains("afm;sdmcesocsdjfe", "ce")
	fmt.Println(b) //                                   true

	//统计字符串中有几个字符串 strings.Count
	c := strings.Count("numberrrrrrrrerererr", "er")
	fmt.Println(c) //                                      4

	//字符串比较 strings.EqualFold
	d := strings.EqualFold("abc", "AbC") //不区分大小写     //true
	fmt.Println(d)
	fmt.Println("abc" == "AbC") //区分大小写                //false

	//返回字符串第一次出现的Index值，没有就-1
	e := strings.Index("Agdsgsmld;dsggsgsgs", "gs") //4
	fmt.Println(e)

	//最后一次出现的
	f := strings.LastIndex("Agdsgsmld;dsggsgsgs", "gs")
	fmt.Println(f)

	//子串替换       strings.Replace(str,a,b,n)//n=-1全替换
	h := "GoGoGo"
	h = strings.Replace(h, "Go", "北京", 2)
	fmt.Println(h)

	//字符串拆分
	strArr := strings.Split("afdsk,sdfs,fds,f", ",") //按,拆分
	fmt.Println(strArr)

	//字符串大小写转换Strings.ToLower//strings.ToUpper
	strtest := "AaAAaAa"
	strtest = strings.ToLower(strtest)
	fmt.Println(strtest)

	//去除字符串两边空格 strings.TrimSpace
	str = strings.TrimSpace("      sss       ")
	fmt.Println(str)

	//去除左右两边叹号和空格********
	str = strings.Trim("!   ! ! sss !  ! !", " !")
	fmt.Println(str)

	//判断字符串是否以指定字符串开头 strings.HasPrefix()//HasSuffix()
	b = strings.HasPrefix("ftp://127.0.0.1", "ftp")
	fmt.Println(b)

	b = strings.HasSuffix("ftp://127.0.0.1", ".0.1")
	fmt.Println(b)
}
