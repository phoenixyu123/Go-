package testingdemo02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	//先序列化再保存到文件
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}
	fp := "C:/study/GoList/src/SGG/chapter15/testingdemo02/test.log"
	//ioutil.WriteFile(filepath,[]byte(data),filemode)
	err = ioutil.WriteFile(fp, data, 0666)
	if err != nil {
		fmt.Println("写入错误", err)
		return false
	}
	return true

}

func (this *Monster) Restore() bool {
	//先从文件读取字符串
	fp := "C:/study/GoList/src/SGG/chapter15/testingdemo02/test.log"
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("读取错误", err)
		return false
	}
	//读取到[]byte data
	// fmt.Println(string(data))
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("反序列化错误!", err)
		return false
	}
	return true
}
