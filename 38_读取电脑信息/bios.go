/*
@ Time : 2020/6/30 10:23
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/
package main

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	getBiosSerialNumber()
}

// 获取bios SerialNumber
func getBiosSerialNumber() string {
	sysType := runtime.GOOS

	if sysType == "linux" {
		// LINUX系统
		log.Println("linux") // 生产中去掉

		cmd := exec.Command("/bin/sh", "-c", `dmidecode -t system`) // linux命令
		serialNo, err := cmd.Output()
		if err != nil {
			log.Println("获取BIOS信息失败")
			return ""
		}
		biosMsg := string(serialNo)
		log.Println("BIOS Msg: ", biosMsg) // 生产中去掉

		start := strings.Index(biosMsg, "Serial Number: ")
		end := strings.Index(biosMsg, "UUID:")
		bios := biosMsg[start+len("Serial Number: ") : end-(len("\n")+1)]
		if len(bios) <= 0 {
			log.Println("获取BIOS Serial Number信息失败")
			return ""
		} else {
			log.Println("BIOS Serial Number: ", bios)
			return bios
		}
	} else if sysType == "windows" {
		// windows系统
		log.Println("windows") // 生产中去掉

		cmd := exec.Command("CMD", "/C", "wmic baseboard get serialnumber")
		serialNumber, err := cmd.Output()
		if err != nil {
			log.Println("获取BIOS信息失败")
			return ""
		}
		serialNumberString := string(serialNumber)
		log.Println("BIOS serialNumber Msg: ", serialNumberString) // 生产中去掉

		start := strings.Index(serialNumberString, "SerialNumber")
		bios := serialNumberString[start+len("Serial Number   \r\r\n") : len(serialNumberString)-len("  \r\r\n\r\r\n")]
		if len(bios) <= 0 {
			log.Println("获取BIOS Serial Number信息失败")
			return ""
		} else {
			log.Println("BIOS Serial Number: ", bios)
			return bios
		}
	} else {
		log.Println("无法识别操作系统")
		return ""
	}
}
