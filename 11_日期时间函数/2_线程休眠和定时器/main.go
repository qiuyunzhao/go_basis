package main

import (
	"fmt"
	"time"
)

func main() {

	time.AfterFunc(time.Second*3, func() {
		fmt.Println("3秒后执行的程序")
	})

	fmt.Println("1")
	time.Sleep(1e9) //单位纳秒 1e9 ns = 1s
	fmt.Println("2")
	time.Sleep(time.Millisecond * 1000) //休眠1s
	fmt.Println("3")
	time.Sleep(time.Second) //休眠1s
	fmt.Println("4")

	//AfterFunc 要想执行主函数不能先结束
	time.Sleep(time.Second * 5)
}
