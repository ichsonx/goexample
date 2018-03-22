package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func main() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

