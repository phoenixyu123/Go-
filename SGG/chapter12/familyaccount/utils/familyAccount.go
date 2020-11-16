package utils

import (
	"fmt"
	"strings"
)

type FamilyAccount struct {
	Key     string
	Loop    bool
	Flag    int
	Balance float64
	Money   float64
	Note    string
	Details string
}

//绑定结构体方法

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		Key:     "",
		Loop:    true,
		Flag:    0,
		Balance: 10000.0,
		Money:   0.0,
		Note:    "",
		Details: "收支\t账户金额\t收支金额\t说		明\n",
	}
}

//1,显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("---------------------------记账---------------------------")
		fmt.Println("                         1.明细")
		fmt.Println("                         2.收入")
		fmt.Println("                         3.支出")
		fmt.Println("                         4.退出")
		fmt.Println("                         输入1-4")
		fmt.Scanln(&this.Key)
		switch this.Key {
		case "1":
			this.Mingxi()
			// fmt.Println("明细记录")
		case "2":
			this.Dengjishouru()
			this.Flag++
		case "3":
			this.Dengjizhichu()
		case "4":
			this.Tuichu()
		default:
			fmt.Println("错误")
			fmt.Println(this.Flag)
			break
		}
		if !this.Loop {
			break
		}
	}

}

func (this *FamilyAccount) Mingxi() {
	if this.Flag == 0 {
		fmt.Println("------------------------当前无记录------------------------")
		return
	}
	//对detail进行拼接处理
	fmt.Println("---------------------当前收支明细记录---------------------")
	// this.Details := "收支\t账户金额\t收支金额\t说		明"
	fmt.Println(this.Details)

}

func (this *FamilyAccount) Dengjishouru() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.Money)
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.Note)
	this.Balance += this.Money
	//拼接
	this.Details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", this.Balance, this.Money, this.Note)
}
func (this *FamilyAccount) Dengjizhichu() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.Money)
	if this.Money > this.Balance {
		fmt.Printf("想支出%v时余额不足	余额：%v\n", this.Money, this.Balance)
		return
	}
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.Note)
	this.Balance -= this.Money
	//拼接
	this.Details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", this.Balance, this.Money, this.Note)
	this.Flag++
}
func (this *FamilyAccount) Tuichu() {
	sure := ""
	fmt.Println("确定？ y/n")
	for {
		fmt.Scanln(&sure)
		if strings.ToLower(sure) == "y" {
			this.Loop = false
			return
		} else if strings.ToLower(sure) == "n" {
			return
		}
		fmt.Println("确定？ y/n")
	}
}
