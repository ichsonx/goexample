package mycommon

import (
	"strings"
	"github.com/axgle/mahonia"
)

func Check(e error)  {
	if e != nil {
		panic(e)
	}
}

//转码函数，根据参数中的charset来转换成utf8。这个函数使用第三方包mahonia
func ConvertString2Utf8(content string, charset string) string  {
	var result string
	if content != ""{
		switch strings.ToLower(charset) {
		case "gbk", "gb18030":
			enc := mahonia.NewDecoder("gbk")
			result = enc.ConvertString(content)
		case "utf8", "utf-8":
			fallthrough
		default:
			result = content
		}
	}
	return result
}
