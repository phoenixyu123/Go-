package main

import "fmt"

func main() {
	var sec float64
	var gender string
	fmt.Scanln(&sec, &gender)
	if sec < 8 {
		if gender == "woman" {
			fmt.Println("女子组决赛")
		} else {
			fmt.Println("男子组决赛")
		}
	} else {
		fmt.Println("淘汰")
	}
}
