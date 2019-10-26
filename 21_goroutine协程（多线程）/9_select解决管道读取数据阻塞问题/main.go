package main

import (
	"fmt"
	"time"
)

//解决必须先关闭管道才能对管道进行遍历的问题，如果不关闭会阻塞而导致deadlock
//使用select可以不用关闭管道，解决从管道取数据的阻塞问题
func main() {

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭可能会阻塞而导致 deadlock

	//问题：在实际开发中，可能我们不好确定什么时候关闭该管道。可以使用select 方式可以解决
	finished := false
	for !finished {
		select {
		//注意: 这里，如果intChan一直没有关闭，不会一直阻塞而deadlock，会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			finished = true
		}
	}

	//1.每个case必须是一个IO操作
	//2.哪个case可以执行就执行哪个
	//3.如果所有case都不能执行，则执行default
	//4.执行完某个case或default会退出
	//5.如果所有case都不能执行，且没有default，子协程中将会阻塞，主协程阻塞会报deadlock

}
