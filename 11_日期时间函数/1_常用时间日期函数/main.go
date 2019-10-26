package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("-------------------------------------------  1  -------------------------------------------------")
	//1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v \n", now)      //值
	fmt.Printf("now type=%T \n", now) //类型

	fmt.Println("-------------------------------------------  2  -------------------------------------------------")
	//2.通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	fmt.Println("-------------------------------------------  3  -------------------------------------------------")
	//格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//方式1：fmt.Sprintf()
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	//方式：Time.Format()
	//每个"2006-01-02 15:04:05"、"2006"、"15:04:05"中的数字是固定的，但可以通过符号来组合生成想要的日期时间格式
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006"))         //当前-年
	fmt.Println(now.Format("01"))           //当前-月
	fmt.Println(now.Format("02"))           //当前-日
	fmt.Println(now.Format("15:04:05"))     //当前-时间
	fmt.Println(now.Format("15"))           //当前-时
	fmt.Println(now.Format("04"))           //当前-分
	fmt.Println(now.Format("05"))           //当前-秒
	fmt.Println(now.Format("2006 15: -05")) //当前-时间

	fmt.Println("-------------------------------------------  4  -------------------------------------------------")
	//Unix和UnixNano的使用（用于获取随机数）
	fmt.Printf("unix时间戳=%v \n", now.Unix())         //秒时间戳
	fmt.Printf("unixnano时间戳=%v \n", now.UnixNano()) //纳秒时间戳

	// 1.通过当前时间戳获取seed种子数
	cTimeStamp := time.Now().UnixNano() // 纳秒的时间戳，更精确
	rand.Seed(cTimeStamp)               // 设置获取随机数的种子数,int64数字即可,根据时间的变动来随机seed
	// 2.获取随机数
	var a = rand.Intn(100) // [0,100)内随机数
	fmt.Println(a)

	fmt.Println("-------------------------------------------  5  -------------------------------------------------")
	//需求，每隔1秒中打印一个数字，打印到100时就退出
	//需求2: 每隔0.1秒中打印一个数字，打印到100时就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100) //休眠100ms
		if i == 5 {
			break
		}
	}

}
