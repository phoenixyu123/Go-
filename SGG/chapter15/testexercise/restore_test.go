package main

import "testing"

// type Monster struct {
// 	Name  string
// 	Age   int
// 	Skill string
// }

func TestReStore(t *testing.T) {
	var monster Monster
	monster.Store()
	if monster.ReStore() != "{\"Name\":\"牛魔王\",\"Age\":18,\"Skill\":\"猪突猛进\"}" {
		t.Fatalf("结果错误\n")
	}
	t.Logf("Success!\n")
}
