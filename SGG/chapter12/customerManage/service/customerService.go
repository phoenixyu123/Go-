package service

import (
	"SGG/chapter12/customerManage/model"
)

//完成显示客户列表的功能

//CS完成对Customer的操作包括增删改查*
type CustomerService struct {
	customers []model.Customer //切片**************************非常重要，围绕着切片来操作所有关于用户的方法
	//声明一个字段表示当前切片customers 已经含有多少个客户

	customerNum int //还可以表示新客户的Id = ++customerNum
}

//创建一个CS结构体用于增上改查的操作，
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "male", 20, "112", "zs@123.com") //此处需要理解,调用model包中的工厂模式，创建一个
	//名为张三的用户
	customerService.customers = append(customerService.customers, *customer) //通过字符串拼接append实现
	//切片的附加，注意此时customer在工厂模式创建下返回的是一个指针，因此要加上一个*
	return customerService
}

//返回用户切片,让view层自己去处理显示的问题，service层只进行用户数据的传递
func (this *CustomerService) CustomerList() []model.Customer {
	return this.customers
}

//添加客户到customers切片中
func (this *CustomerService) AddUser(customer model.Customer) bool {

	this.customers = append(this.customers, customer)

	return true
}
