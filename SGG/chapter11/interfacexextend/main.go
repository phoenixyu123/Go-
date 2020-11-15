package main

//220
//继承的真正价值：解决复用性、可维护性
//接口的真正价值：设计*，设计好规范的方法*，让自定义类型去实现这些方法，更灵活，解耦
//继承是is-a，接口是like-a

import "fmt"

type Monkey struct {
	Name string
}
type BirdAble interface {
	flying()
}
type FishAble interface {
	swimming()
}

func (this *Monkey) climbing() {
	fmt.Printf("%v天生会爬树\n", this.Name)
}

//注意此处使用接口方法的意义是，除了wukong其他的Monkey本
//不应该会飞
func (this LittleMonkey) flying() {
	fmt.Printf("%v会飞\n", this.Name)
}
func (this LittleMonkey) swimming() {
	fmt.Printf("%v会游泳\n", this.Name)
}

type LittleMonkey struct {
	Monkey
}

func main() {
	var lm1 LittleMonkey
	lm1 = LittleMonkey{
		Monkey{
			Name: "Wukong",
		},
	}
	lm1.climbing()
	var ba BirdAble = lm1
	ba.flying()
	var fa FishAble = lm1
	fa.swimming()
	lm1.flying()
	lm1.swimming()
}
