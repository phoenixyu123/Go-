package model

type Student struct {
	Name  string
	Age   int
	Score float64
}
type student struct {
	Name  string
	Age   int
	score float64
}

//注意工厂函数首字母必须大写
func NewStudent(n string, a int, s float64) *student { //工厂模式
	return &student{
		Name:  n,
		Age:   a,
		score: s,
	}
}

func (stu *student) GetScore()float64{
	return stu.score //
}
