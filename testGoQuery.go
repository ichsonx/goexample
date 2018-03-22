package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

//真心不好用啊goquery
func main() {
	doc, err := goquery.NewDocument("http://w3.afulyu.info/pw/thread.php?fid=22&page=5")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Html())
	doc.Find("tr[class$=t_one]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//band := s.Find("a").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
		fmt.Println(i)
		fmt.Println(s.Html())
		s.Children()
	})
}

