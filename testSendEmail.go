/*
	2018-03-05
	1.使用三方库来发送email
	2.使用go内置函数smtp发送email（待完善）
*/

package main

import (
	"net/smtp"
	"strings"
	"fmt"
	"gopkg.in/gomail.v2"
)

var (
	pw = "tmakpyqrpqkkbjfa"
	user = "sonxz@qq.com"
	host = "smtp.qq.com"
	to = []string{"104024786@qq.com"}
	port = 465
)

func main() {
	SendEMailByGomail()
}

//使用第三方库gomail，
// 缺点1：版本管理混乱，分别在github和gopkg.in都有代码和文档，而且文档中有的部分方法在引入代码后，根本就没有，证明作者没有版本更新代码
// 缺点2：不允许多个cc或者bcc，已经查看源代码验证了
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
