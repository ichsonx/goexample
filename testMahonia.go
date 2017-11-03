package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/axgle/mahonia"
	"github.com/antchfx/xquery/html"
)

func main() {
	res, err := http.Get("http://www.szpb.gov.cn/xxgk/qt/tzgg/index_1.htm")
	if err != nil {
		fmt.Println("error")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error
	}

	enc := mahonia.NewDecoder("gbk")
	content := string(body)
	content = enc.ConvertString(content)
	//fmt.Println(content)

}

func queryByXquery(url string) {
	doc, err := htmlquery.LoadURL(url)
	checkerr(err)
	println(doc)
}
func checkerr(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}
