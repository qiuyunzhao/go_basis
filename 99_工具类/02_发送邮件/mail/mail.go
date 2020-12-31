package mail

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

//定义邮箱服务器连接信息(最好配置文件读取)
//例：QQ邮箱： POP3服务器-pop.qq.com:995   SMTP服务器 smtp.qq.com:465或587  （SMTP密码）授权码：wvvwxxuoojwtbfih
var mailConn = map[string]string{
	"host":     "smtp.qq.com",       //发送邮件服务器
	"port":     "587",               //端口
	"user":     "1694644469@qq.com", //发件人账号
	"password": "wvvwxxuoojwtbfih",  //密码或SMTP授权码
}

func SendMail(mailTo []string, subject string, body string, attach interface{}) error {

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	//设置邮件内容
	m.SetHeader("From", mailConn["user"]) //发送者邮箱
	m.SetHeader("To", mailTo...)          //发送给多个用户邮箱
	m.SetHeader("Subject", subject)       //设置邮件主题
	m.SetBody("text/html", body)          //设置邮件正文
	if attach != nil {
		attachfile, ok := attach.(string)
		if ok {
			m.Attach(attachfile) //设置邮件附件
		}
	}

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["password"])

	err := d.DialAndSend(m)
	return err
}

//func 1_常用时间日期函数() {
//	message := gomail.NewMessage()
//
//	//邮件内容设置
//	message.SetHeader("From", "1694644469@qq.com")
//	message.SetHeader("To", "1102401880@qq.com")
//	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
//	message.SetHeader("Subject", "Hello qyz!")
//	message.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
//	message.Attach("attach/code-wallpaper-18.png")
//
//	// 发送邮件服务器、端口、发件人账号、发件人密码
//	// 本QQ邮箱： POP3服务器-pop.qq.com:995   SMTP服务器 smtp.qq.com:465或587  （SMTP密码）授权码：wvvwxxuoojwtbfih
//	d := gomail.NewDialer("smtp.qq.com", 587, "1694644469@qq.com", "wvvwxxuoojwtbfih")
//
//	// Send the email to xxx
//	if err := d.DialAndSend(message); err != nil {
//		log.Println(err)
//	}
//
//	fmt.Println("结束")
//}
