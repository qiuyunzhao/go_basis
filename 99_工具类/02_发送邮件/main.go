package main

import (
	"fmt"
	mail2 "go_basis/99_工具类/02_发送邮件/mail"
	"log"
)

func main() {

	//定义收件人
	mailTo := []string{
		"110XX401880@qq.com",
	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	// 支持基础html
	body := `<p><h3><b>设备报警信息：<b></h3></p>
		<p><b>所属分组:</b> &nbsp` + "DeviceGroup" + `</p>
		<p><b>所属设备:</b> &nbsp` + "DeviceName" + `</p>
        <p><b>报警策略:</b> &nbsp` + "StrategyName" + `</p>
		<p><b>报警原因:</b> &nbsp参数` + "AlarmReason" + `</p>
        <p><b>报警级别:</b> &nbsp` + "AlarmClass" + `</p>
		<p><b>开始时间:</b> &nbsp` + "2020-10-09 11:09:34" + `</p>`
	body = "Good"

	//邮件附件
	attach := "02_发送邮件/attach/code-wallpaper-18.png"

	if err := mail2.SendMail(mailTo, subject, body, attach); err != nil {
		log.Println(err)
	} else {
		fmt.Println("邮件发送成功！")
	}

}
