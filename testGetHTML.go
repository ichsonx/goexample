package main

import (
	"fmt"
	p "github.com/benbjohnson/phantomjs"
	"os"
	"os/exec"
	"log"
	"time"
)

var (
	urls []string
)

func init() {
	for i := 2; i < 100; i++{
		urls = append(urls, fmt.Sprintf("http://w3.afulyu.info/pw/thread.php?fid=22&page=", string(i)))
		urls = append(urls, fmt.Sprintf("http://ac88.info/bt/thread.php?fid=4&page=", string(i)))
		urls = append(urls, fmt.Sprintf("http://w3.afulyu.info/pw/thread.php?fid=5&page=", string(i)))
		urls = append(urls, fmt.Sprintf("http://ac88.info/bt/thread.php?fid=16&page=", string(i)))
	}
}

func main() {
	usePhantomjs()
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
	start := time.Now()

	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}
	Info := log.New(errFile,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
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

	for _, url := range urls{
		if err := page.Open(url); err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println(page.Content())
			Info.Println(page.Content())
		}
	}
	elapsed  := time.Since(start)
	fmt.Printf("coast time:", elapsed )
}
