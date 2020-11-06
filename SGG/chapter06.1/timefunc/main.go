package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now:= %v      now type:=%T\n", now, now)

	//通过now获取年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%02d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	fmt.Println(now.Format("2006 - 01 - 02 15:04:05"))

	//每隔一秒钟打印一个数字，到10
	i := 0
	for {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
		if i == 100 {
			break
		}
		i++
	}

	//Unix,Unixnano---->随机数用(时间戳)
	i = 0
	for {
		fmt.Println(now.UnixNano())
		time.Sleep(time.Second)
	}

}
