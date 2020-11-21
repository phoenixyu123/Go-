package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"This_Is_New_Name"` //修改客户端收到的标签名字
	Age      int    `json:"This_is_New_Age"`  //用到反射机制
	Birthday string
	// birthday string//报错,小写不可导出到别的包(json)
	Sal   float64
	Skill string
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      100,
		Birthday: "2020-02-02",
		Sal:      1234.56,
		Skill:    "猪突猛进",
	}
	//将monster序列化
	//此处monster是指针类型要加取地址符
	str, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

}

func testMap() {
	//使用map之前要make
	a := make(map[string]interface{})
	a["name"] = "111"
	a["age"] = 30
	a["addr"] = "home"
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes)) //发现map并没有顺序

}

func testSlice() {
	var slice []map[string]interface{} //注意var切片不能直接make，但是：=就可以
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})

	m1["name"] = "tom"
	m1["age"] = 3022
	m1["addr"] = [2]string{"hffff", "sgjio"}
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})

	m2["name"] = "tom"
	m2["age"] = 3022
	m2["addr"] = "hffff"
	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}

func testFloat64() {
	var f1 float64 = 1.2345
	data, err := json.Marshal(f1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data)) //对基本类型序列化无意义，就是转成一个字符串
}

func main() {
	//结构体map切片序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64() //基本数据类型的序列化
}
