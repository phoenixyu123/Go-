package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	fmt.Println(cities)

	for k, v := range cities {
		fmt.Printf("key = %v , val = %v\n", k, v)
	}

	fmt.Println("cities有", len(cities), "对key-value")

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

	for i1, v1 := range stuMap {
		fmt.Printf("%v\n", i1)
		for i2, v2 := range v1 {
			fmt.Printf("	stuMap[%v][%v] = %v\n", i1, i2, v2)
		}
	}
}
