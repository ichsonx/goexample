package main

import (
	"os"
	"fmt"
	"github.com/antchfx/htmlquery"
)

func main() {
	f, err := os.Open(`tmp.html`)
	if err != nil {
		panic(err)
	}
	// Parse HTML document
	doc, err := htmlquery.Parse(f)
	if err != nil{
		panic(err)
	}

	// List all matches nodes with the name `a`.
	for _, n := range htmlquery.Find(doc, "//tr[@class=\"tr3 t_one\"]/td/h3/a") {
		fmt.Printf("%s \n", htmlquery.InnerText(n))
	}
}
