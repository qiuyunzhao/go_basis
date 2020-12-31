package mail

import (
	"gopkg.in/gomail.v2"
	"regexp"
	"strconv"
)

//定义邮箱服务器连接信息(最好配置文件读取)
//例：QQ邮箱： POP3服务器-pop.qq.com:995   SMTP服务器 smtp.qq.com:465或587  （SMTP密码）授权码：wvvwxxuoojwtbfih
var mailConn = map[string]string{
	"host":     "smtp.qq.com",        //发送邮件服务器
	"port":     "587",                //端口
	"user":     "16946xx4469@qq.com", //发件人账号
	"password": "wvvwxxxxxxjwtbfih",  //密码或SMTP授权码
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

// 电子邮箱验证
func verifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
