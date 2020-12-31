/*
@ Time : 2020/9/15 15:49
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	generateRandByTime()
}

// 使用时间作为种子生成随机数
func generateRandByTime() {
	fmt.Println(time.Now().UnixNano())

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 核心

	for i := 0; i < 10; i++ {
		fmt.Println(r.Intn(100))
	}
}
