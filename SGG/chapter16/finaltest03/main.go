package main

import (
	"SGG/chapter16/finalanswer"
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

var personNum int = 10
var goNum int = 1
var lock sync.Mutex

func putDataToChan(filePath string, dataChan chan Person, putChanOk chan<- bool) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	finalanswer.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < personNum; i++ {
		randNum := rand.Intn(10) + 2 //2保底
		bytes := make([]byte, randNum)
		for j := 0; j < randNum; j++ {
			b := rand.Intn(26) + 65
			bytes[j] = byte(b)
		}
		randNum2 := rand.Intn(20) + 6
		addr := make([]byte, randNum2)
		for j := 0; j < randNum2; j++ {
			b := rand.Intn(26) + 65
			addr[j] = byte(b)
		}
		person := Person{
			Name:    string(bytes),
			Age:     rand.Intn(100),
			Address: string(addr),
		}
		jsonPerson, err := json.Marshal(person)

		dataChan <- person

		finalanswer.CheckError(err)
		if i == personNum-1 {
			writer.WriteString(string(jsonPerson))
		} else {
			writer.WriteString(string(jsonPerson) + "\n")
		}
	}
	writer.Flush()
	putChanOk <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dataChan := make(chan Person, personNum)

	putChanOk := make(chan<- bool, goNum)
	filePath := "C:/study/GoList/src/SGG/chapter16/finaltest03/test.log"

	go putDataToChan(filePath, dataChan, putChanOk)
	for {
		if len(putChanOk) == 1 {
			close(putChanOk)
			close(dataChan)
			break
		}
	}

	for i := 0; i < personNum; i++ {
		fmt.Println(<-dataChan)
	}

}
