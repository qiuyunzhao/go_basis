package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	function1() //expected exactly 5 fields, found 6: [*/5 * * * * ?] 如果报错参看下边方法
	//errResolve()
}

//----------------------------------------------------------------------------------------------------------------------
//参考： https://www.cnblogs.com/liuzhongchao/p/9521897.html
func function1() {
	i := 0
	c := cron.New()

	//AddFunc
	spec := "*/5 * * * * ?"
	EntryID, err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("EntryID", EntryID)
	}

	//AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	c.Start()      //启动计划任务
	defer c.Stop() //关闭着计划任务, 但是不能关闭已经在执行中的任务.

	select {}
}

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

//---------------------------------------------------------------------------------------------------------------------
//参考： https://blog.csdn.net/Guo_Mao_Zhen/article/details/100890192
func errResolve() {
	c := newWithSeconds()

	c.AddFunc("0/3 * * * * ? ", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})

	c.Start()
	defer c.Stop()

	select {}
}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
