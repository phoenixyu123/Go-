package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LOG_DIRECTORY = "./test.log" //日志地址
)

var onlineConns = make(map[string]net.Conn)

var messageQueue = make(chan string, 1000) //通道

var logFile *os.File
var logger *log.Logger

var quitChan = make(chan bool) //退出通道

//CheckError is a function
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//ProcessInfo is a funciton
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024) //字长1024
	// defer conn.Close()
	// for {
	// 	numOfBytes, err := conn.Read(buf)
	// 	if err != nil {
	// 		break
	// 	}
	// 	if numOfBytes != 0 { //说明收到信息
	// 		remoteaddr := conn.RemoteAddr() //获取远程地址(IP)
	// 		fmt.Printf("From: %s \n信息已收到:= %s \n", remoteaddr, string(buf[0:numOfBytes]))
	// 	}
	// }
	defer func(conn net.Conn) { //不只是Close
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConns, addr) //删除之前有的mapping

		conn.Close()
		for i := range onlineConns {
			fmt.Println("在线:" + i)
		}
	}(conn)

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		if numOfBytes != 0 {
			message := string(buf[0:numOfBytes])
			messageQueue <- message //message传入通道
		}
	}
}

func doProcessMessage(message string) { //判断消息逻辑,判断消息到底是普通信息还是某个指令
	contents := strings.Split(message, "#") //从#切开，切开后用contents切片接收
	if len(contents) > 1 {                  //说明有警号分割--》导致切片数量大于一个
		addr := contents[0]                             //看看你来自哪一个ip
		sendMessage := strings.Join(contents[1:], "\n") //后面是发送的信息

		addr = strings.Trim(addr, " ")         //去掉空格
		if conn, ok := onlineConns[addr]; ok { //如果在onlineConns中确实存在这样一个mapping
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("写入错误")
			}

		}
	} else { //说明在执行命令
		contents := strings.Split(message	, "*")
		if strings.ToUpper(contents[1]) == "LIST" {
			var ips string = ""
			for i := range onlineConns {
				ips = ips + "|" + i
			}
			if conn, ok := onlineConns[contents[0]]; ok {
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("写入错误")
				}
			}
		}
	}
}

//ConsumMessage 消费者函数
func ConsumMessage() { //处理消息
	for {
		select { //等待别人发送消息，根据以下逻辑进行判断
		//使用select来监听chan！！！！！！
		case message := <-messageQueue: //通道里有消息了
			doProcessMessage(message)
		case <-quitChan: //有退出的一个信号
			break

		}
	}
}

//聊天服务器
func main() {
	//打开日志os.OpenFile(地址,权限,打开模式)
	logfile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("日志打开/创建失败")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer logfile.Close() //结束进程关闭文件

	//使用go自带的log打开文件,设置日志的每行开头内容
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	//监听创建    net.Listen(协议,socket)
	listensocket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listensocket.Close() //用defer关闭连接
	fmt.Println("服务器监听...")

	logger.Println("消息写入...") //logger是日志文件

	//开启携程
	go ConsumMessage()

	for { //接收连接
		conn, err := listensocket.Accept()

		CheckError(err)

		addr := fmt.Sprintf("%s", conn.RemoteAddr()) //获取远程地址
		onlineConns[addr] = conn                     //地址与通道互相绑定
		logger.Println("用户进入   IP:", addr)
		for i := range onlineConns {
			fmt.Println(i) //打印这个addr
		}
		go ProcessInfo(conn) //处理请求函数
	}
}
