package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (monster *Monster) Store() string {
	monster = &Monster{
		Name:  "牛魔王",
		Age:   18,
		Skill: "猪突猛进",
	}
	fp := "C:/study/GoList/test.txt"
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	data, err := json.Marshal(&monster)
	if err != nil {
		panic(err)
	}
	writer.WriteString(string(data))
	writer.Flush()
	return string(data)

}
func (monster *Monster) ReStore() string {
	fp := "C:/study/GoList/test.txt"
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var str = ""
	for {
		str, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
	}
	fmt.Println(str) //
	var temp Monster
	err = json.Unmarshal([]byte(str), &temp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)
	return str

}

func main() {
	var monster Monster
	monster.Store()
	monster.ReStore()
}
