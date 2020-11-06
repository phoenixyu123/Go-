package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer+recover来处理panic
	defer func() {
		err := recover() //recover是捕获异常的内置函数
		if err != nil {
			fmt.Println(err)
		}
	}() //defer是执行完函数再出栈
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println("res= ", res)
}

//自定义错误*******                     errors.New("错误说明")
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("文件错误")
	}
}

func test02() {
	err := readConf("config.in")
	if err != nil {
		//读取文件发生错误
		panic(err)
	}
	fmt.Println("test02继续执行")

}

func main() {
	//使用defer+recover来处理panic
	// test()
	// for {
	// 	fmt.Println("错了我也要输出")
	// 	time.Sleep(time.Second)
	// }

	//测试自定义error
	test02()
	fmt.Println("main")
}
