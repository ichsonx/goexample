package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	var output []byte
	var err error
	sender := "sonxz@qq.com"
	smtp_server := "smtp.qq.com"
	to := "sonxz@qq.com"
	port := "465"
	pw := "tmakpyqrpqkkbjfa"
	mail_content := "halo test mail..."
	mail_subject := "test subject"
	cmd := exec.Command("python", "testEmail.py",  "-sender", sender, "-to", to, "-server", smtp_server, "-port", port, "-pw", pw, "-content", mail_content, "-subject", mail_subject)
	if output, err = cmd.Output(); err != nil{
		fmt.Print(err)
		os.Exit(0)
	}
	fmt.Println(output)
}
