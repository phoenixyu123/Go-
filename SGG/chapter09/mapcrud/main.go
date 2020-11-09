package main

import "fmt"

func main() {

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	fmt.Println(cities)

	cities["no3"] = "深圳~" //改
	fmt.Println(cities)

	delete(cities, "no1") //删除
	delete(cities, "no4") //并不报错，不操作
	fmt.Println(cities)

	//查找
	_, ok := cities["no2"]
	if ok {
		fmt.Println("有no2")
	} else {
		fmt.Println("无no2")
	}

	//全部清空
	cities = make(map[string]string)
	fmt.Println(cities)

}
