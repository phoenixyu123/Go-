package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string //修改客户端收到的标签名字
	Age      int
	Birthday string
	// birthday string//报错,小写不可导出到别的包(json)
	Sal   float64
	Skill string
}
type Student struct {
	Name string
	Addr string
	Age  int
}

//反序列化成结构体
func unMarshalStruct() {
	str := "[{\"Addr\":[\"hffff\",\"sgjio\"],\"Age\":3022,\"Name\":\"tom\"},{\"Addr\":\"hffff\",\"Age\":3022,\"Name\":\"tom\"}]"
	var stu [2]Student
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)

	str2 := "{\"Name\":\"牛魔王\",\"Age\":100,\"Birthday\":\"2020-02-02\",\"Sal\":1234.56,\"Skill\":\"猪突猛进\"}"
	var mon Monster
	err = json.Unmarshal([]byte(str2), &mon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mon)

}

func unMarshalMap() {
	str := "{\"addr\":\"home\",\"age\":30,\"name\":\"111\"}"

	var map1 map[string]interface{}
	err := json.Unmarshal([]byte(str), &map1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(map1)
}

func unMarshalSlice() {
	str := "[{\"Addr\":[\"hffff\",\"sgjio\"],\"Age\":3022,\"Name\":\"tom\"}," +
		"{\"Addr\":\"hffff\",\"Age\":3022,\"Name\":\"tom\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice) //unmarshal封装了make
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice)

}

func main() {
	unMarshalStruct()
	unMarshalMap()
	unMarshalSlice(	)
}
