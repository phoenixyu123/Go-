package main

import (
	"SGG/SelfImport/File/JudgeFileExist"
	"fmt"
	"io/ioutil"
)

func main() {

	//将kkk.txt导入到C:/study/kkk.txt
	//1.首先将GoList/kkk.txt内容读取到内存
	//2、将读取的内容写入study/kkk.txt

	fp1 := "C:/study/GoList/kkk.txt"
	fp2 := "C:/study/kkk.txt"
	content, err := ioutil.ReadFile(fp1) //content是[]byte
	if err != nil {
		fmt.Println(err)
		return
	}
	judge, err := JudgeFileExist.PathExist(fp2)
	fmt.Println(judge, err)
	//ioutil.WriteFile(fp,[]byte,mode)
	err = ioutil.WriteFile(fp2, content, 0666)
	if err != nil {
		fmt.Println(err)
	}

}
