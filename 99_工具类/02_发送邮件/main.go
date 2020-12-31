package main

import (
	"fmt"
	mail2 "go_basis/99_工具类/02_发送邮件/mail"
	"log"
)

func main() {

	//定义收件人
	mailTo := []string{
		"1102401880@qq.com",
	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "Good"
	//邮件附件
	attach := "02_发送邮件/attach/code-wallpaper-18.png"

	if err := mail2.SendMail(mailTo, subject, body, attach); err != nil {
		log.Println(err)
	} else {
		fmt.Println("邮件发送成功！")
	}

}
