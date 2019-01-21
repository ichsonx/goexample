package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func main() {
	//m := gomail.NewMessage()
	//m.SetHeader("From", "104024786@qq.com")
	//m.SetHeader("To", "104024786@qq.com")
	////m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	//m.SetHeader("Subject", "Hello!")
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")
	//
	//d := gomail.NewDialer("smtp.qq.com", 587, "104024786@qq.com", "jiuotxwsrtysbhgg")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//
	//// Send the email to Bob, Cora and Dan.
	//if err := d.DialAndSend(m); err != nil {
	//	panic(err)
	//}
	gm()
}

func gm()  {
	m := gomail.NewMessage()
	m.SetHeader("From", "sonxzjw@gmail.com")
	m.SetHeader("To", "sonxzjw@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "sonxzjw@gmail.com", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
