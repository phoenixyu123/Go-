package testingdemo02

import "testing"

func TestStore(t *testing.T) {
	//先创建一个monster实例
	monster := &Monster{
		Name:  "我妻善逸",
		Age:   16,
		Skill: "雷之呼吸",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("测试错误! 期待%v   结果%v\n", true, res)
	}
	t.Logf("测试成功!\n")
}

func TestRestore(t *testing.T) {
	var monster Monster
	res := monster.Restore()
	if !res {
		t.Fatalf("测试错误!\n")
	}
	if monster.Name != "我妻善逸" || monster.Age != 16 || monster.Skill != "雷之呼吸" {
		t.Fatalf("测试错误!\n")
	}
	t.Logf("Success!!!!!!")
}
