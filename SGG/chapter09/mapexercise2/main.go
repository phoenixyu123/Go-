package main

import "fmt"

func main() {
	users := make(map[string]map[string]string)
	users["lisa"] = make(map[string]string)
	users["lisa"]["password"] = "123"
	modifyUser(users, "tom")
	modifyUser(users, "marry")
	modifyUser(users, "lisa")

	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["password"] = "88888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["password"] = "pwd"
		users[name]["nickname"] = "昵称~" + name
	}
}
