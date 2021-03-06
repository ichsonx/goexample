/*
	2018-03-05
	1.使用三方库来发送email（包括了【gomail】【https://github.com/admpub/mail】）
	2.使用go内置函数smtp发送email（待完善）
*/

package main

import (
	"net/smtp"
	"strings"
	"fmt"
	"gopkg.in/gomail.v2"
	"github.com/admpub/mail"
	"github.com/jordan-wright/email"
)

var (
	pw = "tmakpyqrpqkkbjfa"
	user = "sonxz@qq.com"
	host = "smtp.qq.com"
	to = []string{"104024786@qq.com", "sonxz@qq.com"}
	port = 465
)

func main() {
	//SendEMailByAdmPub()
}

//有使用tls的方法，但没有例子，其中的tls.config一大段解说没例子。虽然有bcc、cc。但连ssl都没有邮件都发不了
func SendEMailBymail()  {
	e := email.NewEmail()
	e.From = "Jordan Wright <sonxz@qq.com>"
	e.To = []string{user}
	e.Bcc = []string{"104024786@qq.com"}
	e.Cc = []string{"zjw@cerx.cn"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	e.Send("smtp.qq.com:465", smtp.PlainAuth("", user, pw, "smtp.qq.com"))
}

//使用https://github.com/admpub/mail的第三方
//缺点1：bcc实际效果是“抄送”，而且没有“暗抄”功能
func SendEMailByAdmPub()  {
	conf := &mail.SMTPConfig{
		Username: user,
		Password: pw,
		Host:     host,
		Port:     port,
		Secure:   "SSL",
	}
	c := mail.NewSMTPClient(conf)
	m := mail.NewMail()
	m.AddTo(user) //或 "老弟 <hello@admpub.com>"
	m.AddFrom(user) //或 "老哥 <hank@admpub.com>"
	m.AddSubject("Testing")
	m.AddText("Some text :)")
	m.Bcc = to
	//filepath, _ := os.Getwd()
	//m.AddAttachment(filepath + "/mail.go")
	if e := c.Send(m); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("发送成功")
	}
}

//使用第三方库gomail，
// 缺点1：版本管理混乱，分别在github和gopkg.in都有代码和文档，而且文档中有的部分方法在引入代码后，根本就没有，证明作者没有版本更新代码
// 缺点2：不允许多个cc或者bcc，已经查看源代码验证了
// 不明白网上那么多人使用的原因
func SendEMailByGomail()  {
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", "sonxz@qq.com")
	//查看了源码，无论是cc还是bcc，都只允许1个地址
	//m.SetAddressHeader("Cc", "aaa@qq.com", "")
	//m.SetHeader("Cc", m.FormatAddress("<sonxz@qq.com>, <aaa@qq.com>", "收件人")) //抄送
	m.SetHeader("Bcc",m.FormatAddress("<sonxz@qq.com>, <aaa@qq.com>", "收件人")) //暗送
	m.SetHeader("Subject", "test!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer(host, port, user, pw)
	d.SSL = true

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

//使用golang的原生库smtp来发送邮件
//没有使用ssl/tsl协议发送，即非安全方法，这里使用gmail或者qq邮箱肯定不行
func SendEMailBySMTP()  {
	//配置smtp服务器的身份权限认证
	auth := smtp.PlainAuth("", user, pw, host)
	nickname := "test"
	subject := "test mail"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:465", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
