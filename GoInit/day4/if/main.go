package main

import "fmt"

func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("no")
	// }

	// //else if
	// if age > 35 {
	// 	fmt.Println(">35")
	// } else if age > 18 && age < 34 {
	// 	fmt.Println(">18<34")
	// } else {
	// 	fmt.Println("<18")
	// }

	if age := 19; age > 18 {
		fmt.Println("1")
		//fmt.Println(age)//19
	} else {
		fmt.Println("2")
	}
	//fmt.Println(age) //作用域****，上if中age作用域仅限于if语句中**

}
