package main

import "fmt"

type USB interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

//让Phone实现USB接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

//computer接收接口类型变量
//只要是实现了USB接口方法（start、stop）的
func (c Computer) Working(usb USB) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)

}
