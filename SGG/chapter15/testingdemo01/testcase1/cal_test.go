package main

import (
	"fmt"
	"testing"
)

//编写测试用例,开头必须为Test**********,使用*testing.T 单元测试
//注意!!!!!!!!!测试用例文件名必须!!!以_test结尾,前面无所谓,执行的时候包内所有测试文件一并执行
//调用: go test -v
//将test函数引入testing框架,框架内包含main函数,实际上是testing 进行了import cal_test.go
//然后引用Test开头的函数,因此函数开头必须为Test,而参数必须包含t *testting.T
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		// fmt.Printf("执行错误结果:%v 	期望值:%v\n", res, 55)
		t.Fatalf("AddUpper(10)执行错误结果:%v 	期望值:%v\n", res, 55)
	}
	//如果正确输出日志
	t.Logf("AddUpper(10)执行正确")

}

//go test -v -test.run TestHello可以实现单个函数测试
//go test -v cal_test.go cal.go可以实现单个文件测试
func TestHello(t *testing.T) {
	fmt.Println("Hello函数被调用")
}
