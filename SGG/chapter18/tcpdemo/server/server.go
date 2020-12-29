package main

import (
	"fmt"
	"net"
	"strings"
)

var connMap = make(map[string]net.Conn)

func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	addr := fmt.Sprintf("%s", conn.RemoteAddr())
	defer delete(connMap, addr)
	fmt.Printf("等待客户端ip:%v 的Write\n", conn.RemoteAddr())
	for {
		//创建新切片用于conn.Read()的输入
		buf := make([]byte, 1024)
		ip := conn.RemoteAddr()
		//等待client通过conn发送消息
		//如果没有write，则阻塞

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		//显示Client发送的内容到Server终端
		fmt.Print(ip, " : ", string(buf[:n]), "\n") //n表示发送信息的长度
		if strings.ToUpper(string(buf[:n])) == "LIST" {
			for k, v := range connMap {
				msg := k + "\n"
				if v.RemoteAddr() == conn.RemoteAddr() {
					msg = k + "  <-------YOUR ADDRESS\n"
				}
				numOfBytes, err := conn.Write([]byte(msg))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("向客户端:%v 发送了%vByte的数据:%v\n", conn.RemoteAddr(),
					numOfBytes, k)
			}
		}

	}
}

func main() {
	fmt.Println("监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("error")
	}

	defer listen.Close()
	//等客户端连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err")
		} else {
			ip := fmt.Sprintf("%s", conn.RemoteAddr())
			connMap[ip] = conn
			fmt.Printf("Accept() success con=%v,客户端ip= %v\n",
				conn, conn.RemoteAddr())
		}

		//开启协程为客户端服务
		go process(conn)
	}

	fmt.Printf("listen suc= %v\n", listen)

}
