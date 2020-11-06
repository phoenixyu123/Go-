package main

import "fmt"

func main() {
	heros := [...]string{"宋江", "及时雨", "吴用"}
	for i, v := range heros {
		fmt.Println(i, v)
	}
}
