package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"gopkg.in/gomail.v2"
)

type Subject string

// 预定义主题，也就是邮件的标题
const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

func send(name, subject, body string) error {
	e := global.Config.Email
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) //谁发的
	m.SetHeader("To", mailTo)                                //发给谁
	m.SetHeader("Subject", subject)                          //主题
	m.SetBody("text/html", body)                             //内容
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}

func main() {
	core.InitConfig()
	core.InitLogger()
	NewCode().Send("fckyuo@qq.com", "验证码是:2056")
}
