package main

import "fmt"

func main() {

	var a map[string]string
	// a["num1"] = "haha"//没有数据空间panic: assignment to entry in nil map
	//使用前先make
	a = make(map[string]string)
	a["num1"] = "haha"
	a["num2"] = "haha"
	a["num1"] = "nono" //进行覆盖
	a["num0"] = "ex"
	fmt.Println(a)

	//第二种
	cities := make(map[string]string)
	cities["n1"] = "北京"
	cities["n2"] = "天津"
	cities["n3"] = "上海"
	fmt.Println(cities)

	//第三种
	var heroes map[string]string = map[string]string{
		"no1": "宋江",
		"no2": "及时雨",
	}
	fmt.Println(heroes)

	//map[string]map[string]string
	stuMap := make(map[string]map[string]string)
	stuMap["no1"] = make(map[string]string)
	stuMap["no1"]["name"] = "xiaom"
	stuMap["no1"]["sex"] = "male"
	stuMap["no1"]["addr"] = "beijing"
	stuMap["no2"] = make(map[string]string)
	stuMap["no2"]["name"] = "xiaoh"
	stuMap["no2"]["sex"] = "female"
	stuMap["no2"]["addr"] = "shanghai"
	fmt.Println(stuMap)

}
