/*
	2018-03-01
	支持中文乱码编码解码的第三方包，google也提供了个原生包golang.org/x/text/encoding/simplifiedchinese（https://github.com/golang/text/tree/master/encoding/simplifiedchinese,墙的原因，依赖很难下载，算了）
*/

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/axgle/mahonia"
	"github.com/antchfx/xquery/html"
	"strings"
)

func main() {
	res, err := http.Get("http://www.szpb.gov.cn/xxgk/qt/tzgg/index_1.htm")
	if err != nil {
		fmt.Println("error")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkerr(err)

	enc := mahonia.NewDecoder("gbk")
	content := string(body)
	content = enc.ConvertString(content)
	fmt.Println(content)
}

func QueryByXquery(url string) {
	doc, err := htmlquery.LoadURL(url)
	checkerr(err)
	println(doc)
}
func checkerr(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}


