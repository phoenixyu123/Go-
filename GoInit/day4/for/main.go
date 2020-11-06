package main

import "fmt"

func main() {
	//基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i:=", i) //0-9*
	// }

	// //变1
	// a := 5
	// for ; a < 10; a++ {
	// 	fmt.Println("a:=", a) //5-9*
	// }

	//变2
	// for i := 5; i < 10; {
	// 	i++
	// 	fmt.Println(i) //6-10
	// }
	// i := 1000
	// for {
	// 	fmt.Println("123")
	// 	i--
	// 	if i < 0 {
	// 		break
	// 	}
	// }

	//*************for range 循环*****************
	s := "hello沙河"
	for i, v := range s {
		// fmt.Println(i, v)
		fmt.Printf("%d%c\n", i, v)
	}
}
