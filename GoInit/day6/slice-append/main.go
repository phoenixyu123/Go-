package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳", "天津"}
	fmt.Printf("cap(s1)=%d\n", cap(s1))
	//s1[3] = "广州"  错误，index out of range 索引越界

	//调用append扩充必须用原来的变量接收
	s1 = append(s1, "广州")
	fmt.Printf("cap(s1)=%d\n", cap(s1)) //此处体积是翻倍了，而不是用多少加多少

	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%s cap(s1)=%d\n", s1, cap(s1))

	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) //...表示拆开
	fmt.Println(s1, cap(s1))
}
