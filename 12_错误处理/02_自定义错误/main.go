/*
@ Time : 2020/12/31 10:14
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/

package main

import (
	"errors"
	"fmt"
	cockroachdb_errors "github.com/cockroachdb/errors"
)

func main() {
	fmt.Println("---------------------------- Go语言标准库 ----------------------------------")
	testSimpleCustomError()

	fmt.Println("------------------------- CockroachDB 新建错误 -----------------------------")
	testCockroachdbError()

	fmt.Println("------------------ CockroachDB 添加报错信息上下文前缀 -----------------------")
	testCockroachdbErrorWrap()

	fmt.Println("-------------------------- CockroachDB 次要错误 ----------------------------")
	testCockroachdbSecondaryError()
}

// =================================================== Go语言标准库 ====================================================

// Go 标准库 errors
// 使用 fmt.Errorf() 构造 Go 中最常见的“简单”错误对象类似于带有错误接口的包含在结构中的字符串：打印错误对象只会显示该字符串。
// 使用 Go 的错误包 errors 的构造函数构建错误 errors.New() 结果一样。

func testSimpleCustomError() {
	err := fmt.Errorf("fmt.Errorf")
	fmt.Println(err) // "hello"

	err = errors.New("errors.New")
	fmt.Println(err) // "errors.New"
}

// ================================================ CockroachDB 错误库 =================================================

// ----------------------------------------------- 1. CockroachDB 新建错误 ----------------------------------------------
//
// 使用 Dave Cheney 的错误库[1]，或者甚至更好的 CockroachDB 错误库[2]（通过导入 github.com/cockroachdb/errors）
// 则简单错误也会在构造错误时自动捕获堆栈跟踪。
//
// 建议使用CockroachDB库的以下方法：
// errors.New()：直接替换 Go 标准库的 errors.New()，但它会带有堆栈跟踪；
// errors.Errorf() 或 errors.Newf()：用堆栈跟踪的方式替换 Go 标准库的 fmt.Errorf()；

func testCockroachdbError() {

	err := cockroachdb_errors.New("cockroachdb_errors.New")
	fmt.Println("只会显示错误字符串：", err)

	fmt.Println("显示错误堆栈信息：", err)
	fmt.Printf("%+v \n", err) // 推荐
}

// ----------------------------------------- 2. CockroachDB 添加错误上下文前缀 ------------------------------------------
//
// 当从多个位置调用相同的逻辑，并且可能因错误而失败时，则希望将消息前缀添加到任何返回的错误对象。
// 这有助于提供有关“错误发生的位置”的更多上下文，以便在运行时出现错误时（何时出现错误），可以清楚地了解哪个代码路径产生了错误。
// 在 channel 的场景出现错误是特别有用。
func foo() error { return errors.New("boo") }

func bar() error {
	if err := foo(); err != nil {
		return cockroachdb_errors.Wrap(err, "bar")
	}
	return nil
}

func baz() error {
	// 当提供 nil 错误作为输入时，errors.Wrap() 返回nil。这使我们可以消除 if err != nil 条件。
	return cockroachdb_errors.Wrap(foo(), "baz")
}

func testCockroachdbErrorWrap() {
	err1 := bar()
	fmt.Printf("%+v \n", err1)
	err2 := baz()
	fmt.Printf("%+v \n", err2)
}

// ----------------------------------------- 3. CockroachDB 次要错误 --------------------------------------------
//
// 如果在处理错误时遇到错误，该怎么办？
// 我们希望以某种方式返回有关这两个错误的详细信息，以帮助进行故障排除。
// 同时，出于原因分析的目的，我们要谨慎地将遇到的第一个错误保留为 “主要” 错误。
// 次要错误注解不会影响主要错误上返回的文本，代码的行为就像仅发生了主要错误一样。但是，在详细打印过程中会显示第二个错误的信息;

func testCockroachdbSecondaryError() {
	err := errors.New("主要错误")
	err = cockroachdb_errors.Wrap(err, "主要信息前缀")
	err = cockroachdb_errors.WithSecondaryError(err, cockroachdb_errors.New("次要错误"))

	fmt.Println(err) // 只打印 "主要错误"

	fmt.Printf("%+v \n", err) // 打印只要错误和次要错误的堆栈信息
}
