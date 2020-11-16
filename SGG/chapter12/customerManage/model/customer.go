package model

import "fmt"

//声明customer结构体表示客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//定义工厂模式
func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//定义菜单
func (customer *Customer) MainMenu() {
	for {

	}
}

//返回用户的信息
func (this *Customer) GetInfo() string {
	info := fmt.Sprintf("\t%v\t%v\t%v\t%v\t%v\t%v\n", this.Id, this.Name, this.Gender,
		this.Age, this.Phone, this.Email)//格式化后的字符串
	return info
}
