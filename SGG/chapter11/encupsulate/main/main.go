package main

import (
	"SGG/chapter11/encupsulate/model"
	"fmt"
)

func main() {

	p := model.NewPerson("Smith")
	// p.age = 18// error
	p.SetAge(18)
	fmt.Println(*p)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())

}
