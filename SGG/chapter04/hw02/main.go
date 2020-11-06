// 定义一个变量保存华氏温度，华氏温度转化为摄氏温度:5/9*(华氏温度-100)，求出对应摄氏温度
package main

import "fmt"

func main() {
	var Hs float32 = 0
	fmt.Scanln(&Hs)
	fmt.Printf("%f华氏度对应的摄氏度为:%f摄氏度\n", Hs, (5.0/9)*(Hs-100))
}
