package __测试用例

import (
	"fmt"
	"testing" //引入go 的testing框架包
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain函数可以在测试函数执行之前做一些其他操作")
	//TestMain（）函数中必须要加 m.Run() 方法，否则测试文件下的测试函数不会被执行
	m.Run()
}

//编写要给测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T) {

	//调用
	res := addUpper(10)
	t.Log("AddUpper(10) 执行结果：", res)

}
