package main

import (
	p "github.com/benbjohnson/phantomjs"
	"fmt"
	"os"
)

func main() {
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

	if err := page.Open("https://google.com"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(page.Content())
}
