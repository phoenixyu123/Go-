package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func cheakError(err error){
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
}



func main(){
	conn,err := net.Dial("tcp","127.0.0.1:1234")
	cheakError(err)
	defer conn.Close()
	go MessageSend(conn)//创建协程
	println("客户端已启动")
	buf:=make([]byte,1024)
	for{//此处用于接收消息
		numOfBytes,err:=conn.Read(buf)
		cheakError(err)
		fmt.Println("收到消息:" + string(buf[:numOfBytes]))
	}
}

func MessageSend(conn net.Conn) {
	var input string
	for{
		reader := bufio.NewReader(os.Stdin)
		data,_,_:=reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "EXIT"{
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(input)) //向通道传入比特
		if err !=nil{
			fmt.Println(err)
			break
		}

	}
}
