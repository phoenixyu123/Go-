package main

import "fmt"

type Person struct {
	Name string
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}
func (p Person) test03() {
	p.Name = "Jack"
	fmt.Println("test03() = ", p.Name)
}

func (p *Person) test04() {
	p.Name = "Mary"
	fmt.Println("test04() = ", p.Name)
}

func main() {
	p := Person{"Tom"}
	test01(p)
	test02(&p)
	(&p).test03() //并不会改变真正p内的值，因为方法并没有接收地址
	fmt.Println("经过test03传过去&p后的p.Name = ", p.Name)
	p.test03()

	p.test04()
	(&p).test04()

}
