/*
	2018-02-28
	获取windows、linux下进程列表信息（linux待补全）
*/
package main

import (
	"os/exec"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(CountWinTask("chrome.exe", GetWinTaskList()))

}

//获取windows下进程列表，返回的是一整个字符串（目前只在win10下测试通过）
func GetWinTaskList() string  {
	var output []byte
	var err error
	// 格式化获得进程列表，还有LIST, TABLE
	cmd := exec.Command("tasklist", "/FO",  "CSV")
	if output, err = cmd.Output(); err != nil{
		fmt.Print(err)
		os.Exit(0)
	}
	liststring := string(output)
	return liststring
}

//根据名字返回有多少个同名进程，一个浏览器多个tab都等于多个进程；java、tomcat等的貌似都叫java.exe
func CountWinTask(taskName string, taskList string) int  {
	var result = 0
	counter := 0
	// 用空格分隔字符串，多个连续空格当作一个处理
	for _, line := range strings.Split(taskList, "\n"){
		//因为windows开始的2行都是些title
		if counter > 1{
			tmp := strings.Split(line, ",")
			//fmt.Println(string([]rune(tmp[0])[1:len(tmp[0])-1]))
			//截取出来的数组中的每个元素都带有双引号，所以做字符串截取处理
			if strings.Trim(tmp[0], " ") != "" && string([]rune(tmp[0])[1:len(tmp[0])-1]) == taskName{
				result ++
			}
		}
		counter ++
	}
	return result
}
