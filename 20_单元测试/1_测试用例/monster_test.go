package __测试用例

import (
	"testing"
)

/*
测试用例使用注意事项
	1.测试用例文件名必须以 _test.go 结尾，例如 monster_test.go
    2.测试函数必须以 Test 开头，一般为 Test+函数名 ，例如 TestStore(t *testing.T)
    3.TestStore(t *testing.T) 形参必须为 *testing.T
    4.一个测试文件中，可以有多个测试函数

    5.t.Logf方法可以输出响应日志
    6.出现错误，可以使用 t.Fatalf 来格式化输出错误信息，并退出测试程序（不会影响main函数执行）

    7.运行测试用例指令 （需要先进入到 _test.go文件所在的文件夹）：
        (1) cmd > go test     [运行正确，无日志； 运行错误，输出日志]
        (2) cmd > go test -v  [运行正确或错误，都输出日志]
    8.测试单个文件中的测试方法 命令语句要带上被测试的原文件：
        cmd > go test -v monster_test.go monstert.go
    9.测试单个方法：
        cmd > go test -v -test.run TestStore

    10.测试用例函数并没有放到 1_常用时间日期函数 函数中，也执行了；原理是运行了Testing框架会将测试函数动态添加到mian函数中
    11.PASS表示测试用例运行成功，FAIL表示测试用例运行失败
*/

//测试用例,测试序列化并存储到文件的 Store 方法
func TestStore(t *testing.T) {

	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火.",
	}
	res := monster.Store()

	if !res {
		t.Fatalf("monster.Store() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功!")
}

func TestReStore(t *testing.T) {

	//先创建一个空的 Monster 实例
	var monster = &Monster{}
	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", "红孩儿", monster.Name)
	}

	t.Logf("monster.ReStore() 测试成功!")
}
