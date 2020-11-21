package main

import (
	"flag"
	"fmt"
)

func main() {

	//定义变量用于接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接受用户命令行中输入的-u后面的参数值
	//"u" 就是-u指定参数
	//"" 默认值
	//说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	//非常重要的操作，转换，必须调用该方法
	flag.Parse()
	//输出
	fmt.Printf("user = %v pwd = %v host = %v port = %v\n", user, pwd, host, port)
}
