package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat = Cat{"小白", 1, "白色"}
	var cat2 = Cat{"咪子", 2, "橘色"}
	var cat3 Cat
	cat3.Name = "小绿"
	cat3.Age = 2
	cat3.Color = "绿色"
	fmt.Printf("%T,%v\n", cat3, cat3)
	fmt.Println(cat1, cat2)

}
