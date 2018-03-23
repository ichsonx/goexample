package main

import (
	"os"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
	"regexp"
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

	// 主要是用htmlquery.find函数来找出节点，每次的操作都是如此。htmlquery中有很多函数可以帮助找到节点下的属性、子节点等内容
	for _, a := range htmlquery.Find(doc, "//tr[@class=\"tr3 t_one\"]/td/h3/a") {
		//这里同样适用htmlquery获取节点的内容，有别于一般的xpath获取内容的方法（一般xpath使用路径方式获取节点文本内容例如：/text()）等等
		title := htmlquery.InnerText(a)
		if strings.Contains(title, "飛鳥") {
			exp := regexp.MustCompile(`[a-zA-Z_]+[-]*\d+`)
			params := exp.FindString(title)
			fmt.Println(title)
			fmt.Println(params)

			//if len(params) <= 0{
			//	fmt.Println("no tag")
			//}else {
			//	for _, p := range params{
			//		fmt.Println(p)
			//		p = strings.Replace(p, "-", "", -1)
			//		fmt.Println(strings.ToLower(p))
			//	}
			//}

			//href := htmlquery.SelectAttr(a, "href")
			//fmt.Println(href)
		}

	}
}
