package main

import (
	"fmt"
	"go_code/35_sendmail/mail"
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
	attach := "35_sendmail/attach/code-wallpaper-18.png"

	if err := mail.SendMail(mailTo, subject, body, attach); err != nil {
		log.Println(err)
	} else {
		fmt.Println("邮件发送成功！")
	}

}
