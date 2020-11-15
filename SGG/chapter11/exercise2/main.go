package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StuSlice []Student

func (ss StuSlice) Len() int {
	return len(ss)
}
func (ss StuSlice) Less(i, j int) bool {
	return ss[i].Score < ss[j].Score
}
func (ss StuSlice) Swap(i, j int) {
	// temp := ss[i]
	// ss[i] = ss[j]
	// ss[j] = temp
	//常用交换方法
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	var ss StuSlice
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("学生%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: 100 * rand.Float64(),
		}
		ss = append(ss, stu)
	}
	for _, v := range ss {
		fmt.Println(v)
	}
	sort.Sort(ss) //系统提供快排
	fmt.Println()
	for _, v := range ss {
		fmt.Println(v)
	}
}
