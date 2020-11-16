package main

///此文件用于客户界面

import (
	"SGG/chapter12/customerManage/model"
	"SGG/chapter12/customerManage/service"
	"fmt"
	"strings"
)

//定义一些主菜单中需要的字段
type CustomerView struct {
	Key  string //用于接收用户的输入
	Loop bool   //用于表示退出
	//增加一个字段customerService
	customerService *service.CustomerService ///////////////////////////////////////////为了去调用service里的各种方法，以及CustomerService结构体中的用户信息切片
	//主要是调用CS中的方法：增删改查
}

//主菜单
func (this *CustomerView) MainMenu() {
	this.Loop = true
	for {
		fmt.Println("-----------------------用户信息管理-----------------------")
		fmt.Println("                      1 添 加 用 户")
		fmt.Println("                      2 修 改 用 户")
		fmt.Println("                      3 删 除 用 户")
		fmt.Println("                      4 用 户 列 表")
		fmt.Println("                      5 退       出")
		fmt.Println("                         输入1-5")
		fmt.Scanln(&this.Key)
		switch this.Key {
		case "1":
			// fmt.Println("添加")
			this.Add()
		case "2":
			// fmt.Println("修改")
			this.Modify()
		case "3":
			// fmt.Println("删除")
			this.Delete()
		case "4":
			// fmt.Println("查看")
			this.List() //查看用户的方法在CS.go中构造,list方法里调用CS中的UserList，调用C中的GetInfo
		case "5":
			// this.Loop = false
			this.Exit()
		default:
			fmt.Println("输入错误")
		}

		if this.Loop == false {
			fmt.Println("已退出")
			break
		}

	}
}

//list
func (this *CustomerView) List() {
	fmt.Println("-----------------------用户信息列表-----------------------")

	//首先，获取到当前所有的客户信息(在切片中)
	customers := this.customerService.CustomerList()
	//显示
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------------列表输出完毕-----------------------")
}

//得到用户输入,添加用户
func (this *CustomerView) Add() {
	fmt.Println("------------------------添加新用户------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	newCustomer := model.Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	} //注意此时没有id，去CS中通过给customerNum+1
	if this.customerService.AddUser(newCustomer) {
		fmt.Println("-----------------------用户添加完毕-----------------------")
	} else {
		fmt.Println("---------------------------失败---------------------------")
	}
	//注意此处id需要系统输入,因为id是唯一的，不能有重复id
}

//删除用户
func (this *CustomerView) Delete() {

	fmt.Println("-------------------------删除用户-------------------------")
	for {
		fmt.Println("请输入删除用户的ID(输入-1退出)：")
		id := 0
		fmt.Scanln(&id)
		if id == -1 {
			return
		}
		for {
			fmt.Println("确认删除？(Y/N):")
			sure := ""
			fmt.Scanln(&sure)
			if strings.ToUpper(sure) == "Y" {
				if this.customerService.DeleteUser(id) { //真正的删除操作，除此以外全为交互代码
					fmt.Println("-------------------------删除成功-------------------------")
				} else {
					fmt.Println("-----------------------找不到该用户------------------------")
				}
				break
			} else if strings.ToUpper(sure) == "N" {
				fmt.Println("-------------------------取消删除-------------------------")
				break
			}
		}

	}

}

//Exit
func (this *CustomerView) Exit() {
	for {
		fmt.Println("确认退出？(Y/N):")
		sure := ""
		fmt.Scanln(&sure)
		if strings.ToUpper(sure) == "Y" {
			this.Loop = false
			return
		} else if strings.ToUpper(sure) == "N" {
			return
		}
	}
}

//Modify
func (this *CustomerView) Modify() {
	fmt.Println("-------------------------修改用户-------------------------")
	for {
		fmt.Println("请输入修改用户的ID(输入-1退出)：")
		id := 0
		fmt.Scanln(&id)
		if id == -1 { //退出
			return
		}
		if this.customerService.FindById(id) == -1 {
			fmt.Printf("找不到编号为 %v 的用户\n", id)
		} else { //存在一个编号为id的用户
			
		}

	}
}

func main() {
	cv := CustomerView{
		Key:  "",
		Loop: true,
	}
	cv.customerService = service.NewCustomerService()

	cv.MainMenu()
}
