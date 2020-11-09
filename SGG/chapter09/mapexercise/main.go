package main

import (
	"fmt"
	"strings"
)

func modifyUser(users map[string]map[string]string, name string) {
	var flag = 0
	var findindex = 0
	for i1, v1 := range users {
		fmt.Printf("users[%v]\n", i1)
		for i2, v2 := range v1 {
			fmt.Printf("	users[%v][%v] = %v\n", i1, i2, v2)
			if v1[i2] == name {
				fmt.Println("FIND IT")
				flag = 1
				findindex = 1
			}
			if flag == 1 && v1[i2] == "password" {
				v1[i2] = "88888888"
			}
		}

		flag = 0
		fmt.Println()
	}
	if findindex == 0 {
		fmt.Println("NOT FOUNT")
		users["no3"] = make(map[string]string)
		users["no3"]["name"] = "nickname"
		users["no3"]["password"] = "pwd"
	}
	findindex = 0
}

func main() {
	account := make(map[string]map[string]string)
	account["no1"] = make(map[string]string)
	account["no1"]["name"] = "xiaom"
	account["no1"]["sex"] = "male"
	account["no1"]["password"] = "beijing"
	account["no2"] = make(map[string]string)
	account["no2"]["name"] = "xiaoh"
	account["no2"]["sex"] = "female"
	account["no2"]["password"] = "shanghai"
	var name string
	for {
		fmt.Scanln(&name)
		if strings.ToUpper(name) == "EXIT" {
			break
		}
		modifyUser(account, name)
		fmt.Println(account)
	}
}
