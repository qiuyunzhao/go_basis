/*
@ Time : 2020/3/23 14:32
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// sync.Once能确保实例化对象的Do方法在多线程环境只运行一次,内部通过互斥锁实现
// sync.Once.Do(f func()),能保证once只执行一次，无论你是否更换once.Do(xx)这里的xx方法,这个sync.Once块只会执行一次。

var once sync.Once

func main() {

	for i := 0; i < 3; i++ {
		go func() {
			once.Do(once1)
			fmt.Println("i=", i)
		}()
		time.Sleep(1000)
	}

	for j := 0; j < 3; j++ {
		once.Do(once2)
		fmt.Println("j=", j)
		//once = sync.Once{} //可以放弃只执行一次
		time.Sleep(1000)
	}

	log.Println("============ end ============")
	time.Sleep(8000)
}

func once1() {
	fmt.Println("--------- once1 ---------")
}

func once2() {
	fmt.Println("--------- once2 ---------")
}
