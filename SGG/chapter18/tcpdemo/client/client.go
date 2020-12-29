package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func messageSend(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		//defer conn.Close()
		data, _, _ := reader.ReadLine()
		input := string(data)
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			return
		}
		numOfBytes, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("向服务器发送了%vByte的消息\n", numOfBytes)
	}
}

func readProcess(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("%v\n", string(buf[:numOfBytes]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Printf("成功 ,conn = %v\n", conn)

	go messageSend(conn)
	//将line发送给服务器,通过conn结构体中的Write函数
	go readProcess(conn)
	for {

	}

	fmt.Printf("客户端退出\n")

}
