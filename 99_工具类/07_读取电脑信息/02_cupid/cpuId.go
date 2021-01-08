/*
@ Time : 2020/6/30 9:28
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/

package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	println("============== windows cpuId1 ====================")

	fmt.Println(getCpuId())

	println("================ windows 磁盘序列号==================")

	cmd := exec.Command("CMD", "/C", "WMIC DISKDRIVE GET SERIALNUMBER")
	serialNo, err := cmd.Output()
	if err != nil {
		return
	}
	fmt.Println(string(serialNo))
}

/**
 * windows 获取电脑CPUId
 */
func getCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
