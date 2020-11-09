package main

//JSON.marshal 序列化

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"` //反引号*
}

func main() {
	var a A
	var b B
	// a = b //ERROR
	a = A(b) //因为字段和个数完全一样**
	fmt.Println(a, b)

	monster := Monster{"牛魔王", 180, "芭蕉扇"}
	fmt.Println(monster)

	//序列化
	jsonStr, err := json.Marshal(monster) //jsonStr : []byte
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonStr))
}
