package main

import (
	"fmt"
	"net"
	"strings"
)

var onlineConn = make(map[string]net.Conn)
var messageQueue = make(chan string,1024)
var quitChain = make(chan bool)


func checkError(err error){
	if err!= nil{
		panic(err)
	}
}

func main(){
	listen,err := net.Listen("tcp","127.0.0.1:1010")
	checkError(err)
	defer listen.Close()
	println("服务端启动")
	go ConsumeMessage()
	for{
		conn,err:=listen.Accept()
		if err!= nil{
			break
		}
		addr:= fmt.Sprintf("%s",conn.RemoteAddr())
		onlineConn[addr] = conn
		for i:= range onlineConn{

			println(i)
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
				doMessageProcess(message)
			case <-quitChain:
				break
		}
	}
}

func doMessageProcess(message string) {
	content := strings.Split(message,"#")
	if len(content)>1{
		addr:= content[0]
		sendMessage:= strings.Join(content[1:],",")
		addr = strings.Trim(addr, " ")         //去掉空格
		if conn,ok := onlineConn[addr];ok{
			buf:=[]byte(sendMessage)
			_,err:=conn.Write(buf)
			checkError(err)
		}else {
			fmt.Println("用户不存在")
		}
	}else {
		content:= strings.Split(message,"*")
		addr := content[0]
		conn,ok:=onlineConn[addr]
		if strings.ToUpper(content[1]) == "LIST"&&ok{
			var ips string = ""
			for i:=range onlineConn{
				ips = ips+"|"+i
			}
			_,err:=conn.Write([]byte(ips))
			checkError(err)

		}

	}


}
