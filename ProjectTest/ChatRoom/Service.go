package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LOGDIRECTOR = "./Log.log" //消息地址
)

var onlineConns = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000) //通道
var quitChain = make(chan bool)
var logFile *os.File
var logger *log.Logger

func checkError(err error) {
	if err != nil {
		println(err)
		panic(err)
	}
}

func doProcessMessage(message string) {
	content := strings.Split(message, "#")
	if len(content) > 1 {
		addr := content[0]
		// Join 将 a 中的子串连接成一个单独的字符串，子串之间用 sep 分隔
		sendMessage := strings.Join(content[1:], "\n")
		// Trim 将删除 s 首尾连续的包含在 cutset 中的字符
		addr = strings.Trim(addr, " ")
		if conn, ok := onlineConns[addr]; ok { // ok bool
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Print("[ERROR]消息写入错误!", err)
			}
		}
		//else {
		//	fmt.Printf("[!]IP:%s不在线\n", addr)
		//}
	} else {
		content := strings.Split(message, "*")
		if strings.ToUpper(content[1]) == "LIST" {
			var ips string = ""
			for i := range onlineConns {
				ips = ips + "|" + i
			}
			if conn, ok := onlineConns[content[0]]; ok {
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("写入错误")
				}
			}

		}

	}
}

func processInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConns, addr)
		conn.Close()
		for i := range onlineConns {
			fmt.Println("在线: " + i)
		}
	}(conn)

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			message := string(buf[:numOfBytes])
			messageQueue <- message
		}
	}
}

func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue: //通道中检测到消息，赋给message
			doProcessMessage(message)
		case <-quitChain:
			break
		}
	}
}

func main() {
	listensocket, err := net.Listen("tcp", "127.0.0.1:1234")
	checkError(err)
	println("服务器监听...")
	defer listensocket.Close()
	//创建协程
	go ConsumeMessage()
	for {
		conn, err := listensocket.Accept()
		checkError(err)
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn
		for i := range onlineConns {
			fmt.Println(i) //打印这个addr
		}
		go processInfo(conn) //处理请求函数
	}
}
