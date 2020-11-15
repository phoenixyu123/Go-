package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个Hero结构体切片类型
type HeroSlice []Hero

//3.实现接口Interface->Less,Len,Swap
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less方法决定你使用什么标准进行排序，如按年龄从小到大排序
//if ( i<j )为真,就进行swap，if( i>=j ),不进行swap,注意，
//此时j=i-6，因此。。。
func (hs HeroSlice) Less(i, j int) bool {
	// return hs[i].Age < hs[j].Age
	return hs[i].Name < hs[j].Name //按名字排序升序
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func bubble(intSlice []int) {
	for i := 0; i < len(intSlice)-1; i++ {
		for j := i; j < len(intSlice)-1-i; j++ {
			if intSlice[j] > intSlice[j+1] {
				temp := intSlice[j]
				intSlice[j] = intSlice[j+1]
				intSlice[j+1] = temp
			}
		}
	}
}



func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对intSlice排序
	//1.冒泡排序
	bubble(intSlice)
	fmt.Println(intSlice)

	//2.使用系统提供的方法
	var intSlice2 = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice2)
	fmt.Println(intSlice2)

	//对结构体切片进行排序
	//1.冒泡
	//2.sort.Sort(data interface)
	//可以接受一个实现了Len()int、Less(i,j int)bool、Swap(i,j int)接口的变量
	var heroes HeroSlice
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(10000)),
			Age:  rand.Intn(100),
		}
		//切片附加
		heroes = append(heroes, hero)
	}
	//排序前顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	sort.Sort(heroes)
	//排序后
	fmt.Println()
	for _, v := range heroes {
		fmt.Println(v)
	}
}
