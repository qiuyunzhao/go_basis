package main

import (
	"fmt"
	"time"
)

//单元测试不会影响函数的执行,运行Testing框架会将测试函数动态添加到mian函数中

func main() {

	i := 0
	for range time.Tick(time.Second * 3) {
		i++
		fmt.Println("mian() is running...", i)
	}

}
