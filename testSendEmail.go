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
)

var (
	pw = "tmakpyqrpqkkbjfa"
	user = "sonxz@qq.com"
	host = "smtp.qq.com"
	to = []string{"104024786@qq.com"}
)

func main() {
	SendEMailBySMTP()
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

//使用golang的原生库smtp来发送邮件,通过tls
//func SendEMailBySMTPByTLS()  {
//	nickname := "test"
//	user := "sonxz@qq.com"
//	subject := "test mail"
//	content_type := "Content-Type: text/plain; charset=UTF-8"
//	body := "This is the email body."
//	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
//		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
//	if err != nil {
//		fmt.Printf("send mail error: %v", err)
//	}
//}
