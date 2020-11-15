package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("姓名：%v 年龄：%v 成绩：%v\n", s.Name, s.Age, s.Score)
}
func (s *Student) SetScore(score int) {
	s.Score = score
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生考试中")
}

func (g *Graduate) testing() {
	fmt.Println("本科生考试中")
}

func main() {
	p1 := &Pupil{}
	p1.Student.Name = "Tom"
	p1.Student.Age = 8
	p1.testing()
	p1.SetScore(80)
	p1.ShowInfo()
	g1 := Graduate{}
	g1.Student.Name = "Mary"
	g1.Student.Age = 100
	g1.testing()
	g1.SetScore(100)
	g1.ShowInfo()
}
