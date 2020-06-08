/*
@ Time : 2020/5/29 14:44
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "18_文件操作/sources/timg.jpg"
	err := os.Remove(file)

	if err != nil {
		// 删除失败
		log.Println(err)
	} else {
		// 删除成功
		fmt.Println("删除成功")
	}
}
