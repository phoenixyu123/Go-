package account

import "fmt"

type account struct {
	UserName string
	password string
	balance  float64
}

func NewAccount(uname string) *account {
	if len(uname) > 5 && len(uname) < 11 {
		return &account{
			UserName: uname,
			balance:  0,
		}
	}
	//fmt.Println("账号长度必须在6-10")
	panic("账号长度必须在6-10")
}

func (a *account) SetPassword(password string) {

	if len(password) != 6 {
		fmt.Println("密码必须为六位")
	} else {
		a.password = password
	}
}
func (a *account) GetPassword() string {
	return a.password
}
func (a *account) SetBalance(balance float64) {

	if balance <= 20 {
		fmt.Println("余额不足20")
	} else {
		a.balance = balance
	}
}
func (a *account) GetBalance() float64 {
	return a.balance
}
