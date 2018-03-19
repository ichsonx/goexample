package main

import (
	"fmt"
	p "github.com/benbjohnson/phantomjs"
	"os"
	"os/exec"
)

func main() {
	useCurl()
}

func useCurl() {
	var output []byte
	var err error
	// 格式化获得进程列表，还有LIST, TABLE
	cmd := exec.Command("curl", "http://w3.afulyu.info/pw/thread.php?fid=22&page=5")
	if output, err = cmd.Output(); err != nil {
		fmt.Print(err)
		os.Exit(0)
	}
	fmt.Println(string(output))
}

func usePhantomjs() {
	if err := p.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer p.DefaultProcess.Close()

	page, err := p.CreateWebPage()
	if err != nil {
		fmt.Println(err)
	}
	defer page.Close()

	if err := page.Open("https://www.google.com"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(page.Content())
	}
}
