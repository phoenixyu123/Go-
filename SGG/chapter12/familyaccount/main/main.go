package main

import "SGG/chapter12/familyaccount/utils"

// type FamilyAccount struct {
// 	Key     string
// 	Loop    bool
// 	Flag    int
// 	Balance float64
// 	Money   float64
// 	Note    string
// 	Details string
// }

func main() {
	utils.NewFamilyAccount().MainMenu()
	//NewFamilyAccount()工厂模式构造方法
}
