package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var personChan chan interface{}
	personChan = make(chan interface{}, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		personChan <- Person{
			Name:    strconv.Itoa(rand.Intn(1000)),
			Age:     rand.Intn(20),
			Address: strconv.Itoa(rand.Intn(1000)),
		}
	}
	for i := 0; i < 10; i++ {
		person := <-personChan
		a := person.(Person)
		fmt.Println(a)
	}

}
