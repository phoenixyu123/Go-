package main

import (
	"fmt"
	"net"
	"strings"
)

var onlineConn = make(map[string]net.Conn)
var quitChain = make(chan bool)
var messageQueue = make(chan string, 1000)

func doMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := strings.Trim(contents[0], " ")

		sendMessage := strings.Join(contents[1:], ",")
		if conn, ok := onlineConn[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			cheakError(err)

		}
	} else {
		if strings.ToUpper(contents[0]) == "LIST" {
			ips := ""
			for i := range onlineConn {
				ips = ips + "|" + i
			}
			_, err := conn.Write([]byte(ips))
		}
	}
}

func processInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s", conn)
		delete(onlineConn, addr)
		conn.Close()

		for i := range onlineConn {
			fmt.Println(i)
		}

	}(conn)
	for {
		numOfBytes, err := conn.Read(buf) //把buf传进去
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			message := string(buf[0:numOfBytes])
			messageQueue <- message
		}

	}
}

func consumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			doMessage(message, conn)
		case <-quitChain:
			break
		}
	}
}

func cheakError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	listensocket, err := net.Listen("tcp", "127.0.0.1:1010")
	cheakError(err)
	defer listensocket.Close()
	fmt.Println("服务器监听")
	go consumeMessage()
	for {
		conn, err := listensocket.Accept()

		cheakError(err)
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConn[addr] = conn
		for i := range onlineConn {
			fmt.Println(i)

		}
		go processInfo(conn)
	}

}
