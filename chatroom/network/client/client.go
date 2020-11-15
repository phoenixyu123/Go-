package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LOG_DIRECTORY = "../test.log" //日志地址
	USER_LOG      = "../log/"
)

var logFile *os.File
var logger *log.Logger
var userlogger *log.Logger

//CheckError 检错
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//MessageSend 通过net.Conn发送数据
func MessageSend(conn net.Conn) {
	var Input string
	for { //创建死循环读取消息
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()       //读取该行比特流
		Input = string(data)                  //比特转字符串
		if strings.ToUpper(Input) == "EXIT" { //ToUpper大写
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(Input))
		if err != nil {
			conn.Close()
			fmt.Println("客户端连接失败" + err.Error())
			break
		}
	}
}

func main() {

	logfile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Printf("打开日志失败\n")
		panic(err)
	}
	defer logfile.Close()

	//使用go自带的log打开文件,设置日志的每行开头内容
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := net.Dial("tcp", "127.0.0.1:21999") //连接一致

	//打开用户个人log
	useraddr := fmt.Sprintf("%s", conn.LocalAddr())
	useraddr = strings.Replace(useraddr, ".", "-", -1)
	useraddr = strings.Replace(useraddr, ":", "-", -1)
	var userLogUrl string = USER_LOG + useraddr + ".log"
	userlog, err := os.OpenFile(userLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Printf("打开用户日志失败\n")
		panic(err)
	}
	defer userlog.Close()
	//使用go自带的log打开文件,设置日志的每行开头内容
	userlogger := log.New(userlog, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	CheckError(err)
	defer conn.Close()
	logger.Println("消息写入...")
	userlogger.Println("消息写入...")
	// conn.Write([]byte("Hello World"))
	// fmt.Println("数据已发送...")
	go MessageSend(conn)      //发送数据, 建立协程MessageSend********<线程<进程
	buf := make([]byte, 1024) //buf 字节切片用于接收数据
	fmt.Println("客户端监听...")
	for { //不断接收消息
		numOfbytes, err := conn.Read(buf)
		CheckError(err)
		fmt.Println("得到消息:" + string(buf[0:numOfbytes])) //防止多余输出
		logger.Printf("\n%s收到消息：\n%s", conn.LocalAddr(), string(buf[0:numOfbytes]))
		userlogger.Printf("\n%s收到消息：\n%s", conn.LocalAddr(), string(buf[0:numOfbytes]))
	}
	// fmt.Println("客户端关闭")
}
