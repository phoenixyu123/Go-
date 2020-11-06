package main

import (
	"fmt"
	"net"
	"strings"
)

var onlineConn = make(map[string]net.Conn)
var messageQueue = make(chan string,1000)
var quitChain = make(chan bool)


func checkError(err error){
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
}

func main(){
	listen,err:= net.Listen("tcp","127.0.0.1:8060")
	checkError(err)
	defer listen.Close()
	fmt.Println("服务器启动")
	go ConsumeMessage()
	for{
		conn,err:= listen.Accept()
		checkError(err)
		addr:=fmt.Sprintf("%s",conn.RemoteAddr())
		onlineConn[addr] = conn
		for i:=range onlineConn{
			fmt.Println(i)
		}
		go processInfo(conn)//容纳用户协程
	}
}

func processInfo(conn net.Conn) {
	buf:=make([]byte,1024)
	defer func(conn net.Conn) {
		addr:=fmt.Sprintf("%s",conn.RemoteAddr())
		delete(onlineConn,addr)
		conn.Close()
		for i:=range onlineConn{
			fmt.Println(i)
		}

	}(conn)

	for{
		numOfBytes,err:=conn.Read(buf)
		if err!=nil{
			break
		}
		if numOfBytes!=0{
			message:=string(buf[:numOfBytes])
			messageQueue<-message
		}
	}
}

func ConsumeMessage() {
	for{
		select {
			case message:=<-messageQueue:
				doMessage(message)
			case <-quitChain:
					break
		}
	}
}

func doMessage(message string) {
	contents:= strings.Split(message,"#")
	if len(contents)>1{
		addr:=contents[0]
		sendMessage:=strings.Join(contents[1:],",")
		if conn,ok := onlineConn[addr];ok{
			_,err:=conn.Write([]byte(sendMessage))
			checkError(err)
		}
	}else {

	}
}
