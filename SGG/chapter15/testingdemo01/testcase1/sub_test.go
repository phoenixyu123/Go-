package main

import (
	"testing"
)

//编写测试用例,开头必须为Test**********,使用*testing.T 单元测试
//调用: go test -v输出详细日志信息如fatalf logf 等,go test 直接显示正确与否其他不显示
//将test函数引入testing框架,框架内包含main函数,实际上是testing 进行了import cal_test.go
//然后引用Test开头的函数,因此函数开头必须为Test,而参数必须包含t *testting.T
//函数名Test后面第一个字母必须是大写!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func TestGetSub(t *testing.T) {
	res := GetSub(10, 5)
	if res != 5 {
		// fmt.Printf("执行错误结果:%v 	期望值:%v\n", res, 55)
		t.Fatalf("GetSub(10,5)执行错误结果:%v 	期望值:%v\n", res, 5)
		//Fatalf格式化输出测试结果并结束执行
	}
	//如果正确输出日志
	t.Logf("GetSub(10,5)执行正确")

}
