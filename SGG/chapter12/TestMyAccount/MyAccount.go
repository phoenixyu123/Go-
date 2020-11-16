package main

import (
	"fmt"
	"strings"
)

func Mingxi(detail string, flag int) {
	if flag == 0 {
		fmt.Println("------------------------当前无记录------------------------")
		return
	}
	//对detail进行拼接处理
	fmt.Println("---------------------当前收支明细记录---------------------")
	// detail := "收支\t账户金额\t收支金额\t说		明"
	fmt.Println(detail)

}

func Dengjishouru(balance *float64, money *float64, note *string, detail *string) {
	fmt.Println("本次收入金额:")
	fmt.Scanln(money)
	fmt.Println("本次收入说明:")
	fmt.Scanln(note)
	*balance += *money
	//拼接
	*detail += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", *balance, *money, *note)
}
func Dengjizhichu(balance *float64, money *float64, note *string, detail *string) {
	fmt.Println("本次支出金额:")
	fmt.Scanln(money)
	if *money > *balance {
		fmt.Printf("想支出%v时余额不足	余额：%v\n", *money, *balance)
		return
	}
	fmt.Println("本次支出说明:")
	fmt.Scanln(note)
	*balance -= *money
	//拼接
	*detail += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", *balance, *money, *note)
}
func Tuichu(loop *bool) {
	sure := ""
	fmt.Println("确定？ y/n")
	for {
		fmt.Scanln(&sure)
		if strings.ToLower(sure) == "y" {
			*loop = false
			return
		} else if strings.ToLower(sure) == "n" {
			return
		}
		fmt.Println("确定？ y/n")
	}
}

func main() {
	key := ""
	loop := true
	balance := 10000.0 //默认余额
	money := 0.0       //收支
	note := ""         //说明
	detail := "收支\t账户金额\t收支金额\t说		明\n"
	flag := 0
	for {
		// if key != "1" && key != "2" && key != "3" && key != "4" {
		fmt.Println("---------------------------记账---------------------------")
		fmt.Println("                         1.明细")
		fmt.Println("                         2.收入")
		fmt.Println("                         3.支出")
		fmt.Println("                         4.退出")
		fmt.Println("                         输入1-4")
		// }

		fmt.Scanln(&key)
		switch key {
		case "1":
			Mingxi(detail, flag)
			// fmt.Println("明细记录")
		case "2":
			Dengjishouru(&balance, &money, &note, &detail)
			flag++
		case "3":
			Dengjizhichu(&balance, &money, &note, &detail)
			flag++
		case "4":
			Tuichu(&loop)
		default:
			fmt.Println("错误")
		}
		if !loop {
			break
		}
	}
	fmt.Println("退出")
}
