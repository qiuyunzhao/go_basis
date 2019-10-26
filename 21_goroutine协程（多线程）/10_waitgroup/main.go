package main

import (
	"fmt"
	"sync"
)

//适用于简单的低水平线程处理（常用于等待所有线程执行完毕）
//原理是计数器
func main() {
	var wg sync.WaitGroup
	wg.Add(5) //添加5个
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("第", i, "此执行")
			wg.Done() //减去1个
		}()
	}

	wg.Wait() //阻塞等待ws减为0
}
