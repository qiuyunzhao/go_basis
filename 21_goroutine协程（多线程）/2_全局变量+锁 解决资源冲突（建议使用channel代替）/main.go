package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：现在要计算 1-20 的各个数的阶乘，并且把各个数的阶乘放入到map中,最后显示出来。

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 定义 map 全局的变量.
// 3. 我们启动的多个协程，统计的将结果放入到 map中

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	// lock 是一个全局的互斥锁，
	// sync 是包: synchornized 同步
	// Mutex : 是互斥
	lock sync.Mutex //互斥锁（锁代码）

	lock1 sync.RWMutex //读写锁（锁数据的读写）
)

// test 函数用于计算 n!, 将结果放入到 myMap
func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//加锁，解决多线程同时操作资源造成的资源争夺问题
	lock.Lock()
	myMap[n] = res //不加锁会报错：concurrent map writes
	//解锁
	lock.Unlock()
}

func main() {

	// 开启多个协程完成任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠5秒钟 - 为了让协程执行完毕（主线程执行完毕，不管协程是否执行完毕，程序会结束）----使用这种方式不合理，时间没法确定
	time.Sleep(time.Second * 2)

	//遍历，输出结果（底层不知道协程是否执行完，不加锁会有问题）
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}
