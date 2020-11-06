package main

import "fmt"

func main() {
	//make
	var slice []float64 = make([]float64, 5, 10)
	//由make创建的切片其底层是由make维护的，对外不可见

	fmt.Println(slice)
	slice[2] = 10
	slice[3] = 20
	fmt.Println("slice:=", slice)
	fmt.Println("slice len:=", len(slice))
	fmt.Println("slice cap:=", cap(slice))

	var strSlice []string = []string{"Tom", "Jack", "Frank"}
	fmt.Println("strSlice:=", strSlice)
	fmt.Println("strSlice len:=", len(strSlice))
	fmt.Println("strSlice cap:=", cap(strSlice))

}
