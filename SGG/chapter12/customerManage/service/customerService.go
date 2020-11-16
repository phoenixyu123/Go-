package service

import (
	"SGG/chapter12/customerManage/model"
	"fmt"
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
func (this *CustomerService) AddUser(customer model.Customer) bool { //！！！此处必须用指针方式进行绑定，否则每次都会分配新的customer
	//此处如果不传递指针，则是值拷贝，会导致每一次都会进行创建一个全新的切片

	//需要确定一个分配id的规则，如添加的顺序cunstemorNum
	this.customerNum++
	customer.Id = this.customerNum //分配ID
	this.customers = append(this.customers, customer)

	return true
}

//根据id查找客户在切片中对应的下标，没有返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1 //默认是没有
	for i, _ := range this.customers {
		if this.customers[i].Id == id {
			index = i
			return index
		}
	}
	return index
}

//根据id删除用户，从切片中
func (this *CustomerService) DeleteUser(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	} else { //---------------------------------------------------------------如何从切片中删除一个元素
		this.customers = append(this.customers[:index], this.customers[index+1:]...) //注意最后要...!!!!!!!!!!!!!因为后面也是一个切片
	}
	return true
}

//根据id获取信息,并进行修改
func (this *CustomerService) ModifyUserInfo(id int) {
	for i, _ := range this.customers {
		if id == this.customers[i].Id {
			fmt.Printf("姓名(%v):", this.customers[i].Name)
			name := ""
			fmt.Scanln(&name)
			if name != "" {
				this.customers[i].Name = name
			}
			fmt.Printf("性别(%v):", this.customers[i].Gender)
			gender := ""
			fmt.Scanln(&gender)
			if gender != "" {
				this.customers[i].Gender = gender
			}
			fmt.Printf("年龄(%v):", this.customers[i].Age)
			age := -1
			fmt.Scanln(&age)
			if age != -1 {
				this.customers[i].Age = age
			}
			fmt.Printf("电话(%v):", this.customers[i].Phone)
			phone := ""
			fmt.Scanln(&phone)
			if phone != "" {
				this.customers[i].Phone = phone
			}
			fmt.Printf("邮箱(%v):", this.customers[i].Email)
			email := ""
			fmt.Scanln(&email)
			if email != "" {
				this.customers[i].Email = email
			}
		}
	}
}
