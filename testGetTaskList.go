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
	list := GetWinTaskList()
	fmt.Println(strings.Count(list, "QQ.exe"))

}

//获取windows下进程列表，返回的是一整个字符串（目前只在win10下测试通过）
func GetWinTaskList() string  {
	var output []byte
	var err error

	cmd := exec.Command("tasklist")
	if output, err = cmd.Output(); err != nil{
		fmt.Print(err)
		os.Exit(1)
	}

	liststring := string(output)
	// 用空格分隔字符串，多个连续空格当作一个处理
	//for _, process := range strings.Split(liststring, "\n"){
	//	fmt.Println(process)
	//}
	return liststring
}

//根据名字返回有多少个同名进程，一个浏览器多个tab都等于多个进程；java、tomcat等的貌似都叫java.exe
func CountWinTask(taskName string) int  {
	return strings.Count(GetWinTaskList(), taskName)
}

//如果taskName这个服务存在，返回真，否则返回假
func WinTaskAlive(taskName string) bool  {
	if CountWinTask(taskName) > 0 {
		return true
	}
	return false
}
