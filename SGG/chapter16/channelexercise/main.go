package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan <- 10
	allChan <- "tom jack"
	allChan <- Cat{
		Name: "咪子",
		Age:  2,
	}
	//希望获得到管道中第三个元素
	<-allChan
	<-allChan
	newcat := <-allChan
	fmt.Printf("newCat类型:%T	newCat值:%v\n", newcat, newcat) //main.Cat,{咪子 2}
	//类型是对的但是不能直接用,应该使用类型断言
	a := newcat.(Cat) //类型断言
	fmt.Println(a.Name)

}
