package main

import "fmt"

//定义结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法
//存款
func (acc *Account) Deposite(money float64, pwd string) {
	if pwd == acc.Pwd {

		if money < 0 {
			fmt.Println("金额小于0")
			return
		}
		acc.Balance += money
		fmt.Println("存款完毕")
	} else {
		fmt.Println("密码错误")
		return
	}
}

//取款
func (acc *Account) WithDraw(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("密码错误")
		return
	}
	if money > acc.Balance {
		fmt.Println("余额不足")
		return
	}
	acc.Balance -= money
	fmt.Println("取款完毕")
}

//查询余额
func (acc *Account) Query(pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("密码错误")
		return
	}
	fmt.Println("余额:= ", acc.Balance)
}

func main() {
	acc := Account{"1111", "666666", 100.00}
	acc.Query("666666")
	acc.Deposite(100, "666666")
	acc.Query("666666")
	acc.WithDraw(9, "666666")
	acc.Query("666666")
}
