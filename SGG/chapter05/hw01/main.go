package main

import "fmt"

func main() {
	//水仙花
	fmt.Println("请输入:")
	var max int = 1000
	// fmt.Scanln(&max)
	for i := 0; i < max; i++ {
		a := i % 10        //个位
		b := (i / 10) % 10 //十位
		c := i / 100       //百位
		if i == a*a*a+b*b*b+c*c*c {
			fmt.Println(i, c, b, a)
		}
	}

}
