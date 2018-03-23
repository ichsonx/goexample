package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
	"os"
	"strings"
)

//真心不好用啊goquery
func main() {
	f, err := os.Open(`tmp.html`)
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Html())
	doc.Find("tr[class$=t_one]").Each(func(i int, s *goquery.Selection) {
		au := strings.Trim(s.Find("a[class=bl]").Text(), " ")
		title := strings.Trim(s.Find("h3 a").Text(), " ")
		url, _ := s.Find("h3 a").Attr("href")
		if au != "" && au != "mh1024"{
			if strings.Contains(title, "超優") {
				fmt.Println(url)
			}
		}
	})
}

