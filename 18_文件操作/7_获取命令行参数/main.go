package main

import (
	"flag"
	"fmt"
	"os"
)

//go build -o text.exe 1_常用时间日期函数.go    (命令行：将 1_常用时间日期函数.go 编译成 text.exe 的可执行文件)

func main() {
	//method1()
	method2()
}

//os.Args获取命令行参数
func method1() {
	//os.Args用来获取命令行参数(以空格间隔)
	fmt.Println("命令行的参数有", len(os.Args))

	//遍历os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}

//flag包解析命令行参数   "text.exe -u root -pwd 123456 -h 127.0.0.5 -port 8080"
func method2() {
	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user             ,就是接收用户命令行中输入的 -u 后面的参数值
	//"u"               ,就是 -u 指定参数
	//""                ,默认值
	//"用户名默认为空"   ,说明
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号默认为3306")

	//这里有一个非常重要的操作,转换， 必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
