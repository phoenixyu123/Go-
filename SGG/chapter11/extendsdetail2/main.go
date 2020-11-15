package main

import "fmt"

type A struct {
	Name  string
	score int
}
type B struct {
	Name string
	age  int
}
type C struct {
	A
	B
}
type D struct {
	a A
}

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}
type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int //匿名字段，调用时直接E.int即可
}

func main() {
	var c C
	// c.Name = "Tom"//error不知道是哪个Name
	c.A.Name = "Tom"
	fmt.Println(c)

	var d D
	// d.Name = "Jack"//error A不再是匿名函数
	d.a.Name = "Jack"
	fmt.Println(d.a.Name)

	tv1 := TV{Goods{"电视机", 5999.88},
		Brand{"杂牌", "北京"}}
	tv2 := TV{
		Goods{
			Name:  "电视剧001",
			Price: 222.22,
		},
		Brand{
			Name:    "HP",
			Address: "BJ",
		},
	}
	tv3 := TV2{&Goods{"电视003", 2299.88},
		&Brand{"创维", "天津"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视剧004",
			Price: 4444.22,
		},
		&Brand{
			Name:    "Lenovo",
			Address: "SH",
		},
	}

	fmt.Println("tv1", tv1)
	fmt.Println("tv2", tv2)
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
	//匿名字段时基本数据类型的使用
	var e E
	e.Name = "Fox"
	e.Age = 300
	e.int = 999
	fmt.Println("e", e)
}
