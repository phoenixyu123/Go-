package main

import "testing"

// type Monster struct {
// 	Name  string
// 	Age   int
// 	Skill string
// }

func TestStore(t *testing.T) {
	var monster Monster
	if monster.Store() != "{\"Name\":\"牛魔王\",\"Age\":18,\"Skill\":\"猪突猛进\"}" {
		t.Fatalf("store输出错误\n")
	}
	t.Logf("输出正确\n")
}
