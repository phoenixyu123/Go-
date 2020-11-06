package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkError(err error){
	if err!=nil{
		panic(err)
	}
}

func main(){
	conn,err:=net.Dial("tcp","127.0.0.1:8060")
	checkError(err)
	defer conn.Close()
	fmt.Println("客户端启动")
	go sendMessage(conn)
	buf:=make([]byte,1024)
	for{
		numOfBytes,err:=conn.Read(buf)
		checkError(err)
		fmt.Println("->",string(buf[:numOfBytes]))
	}
}

func sendMessage(conn net.Conn) {
	var input string
	for{
		reader:=bufio.NewReader(os.Stdin)
		data,_,_:=reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "EXIT" {
			break
		}
		_,err:=conn.Write([]byte(input))
		if err!=nil{
			break
		}
	}
}
