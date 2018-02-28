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
	GetWinTaskList()
}

//获取windows下进程列表，返回的是一个二维slice（目前只在win10下测试通过）
func GetWinTaskList() [][]string {
	var output []byte
	var err error
	var tasklist [][]string
	// 格式化获得进程列表，还有LIST, TABLE
	cmd := exec.Command("tasklist", "/FO",  "CSV", "/V")
	if output, err = cmd.Output(); err != nil{
		fmt.Print(err)
		os.Exit(0)
	}

	counter := 0
	// 用空格分隔字符串，多个连续空格当作一个处理
	for _, line := range strings.Split(string(output), "\n"){
		//因为windows开始的2行都是些title
		if counter > 2{
			tmp := strings.Split(line, ",")
			//白痴的windows，竟然还会有空！？
			if tmp[0] == "" || tmp[1] == "" || tmp[2] == "" || tmp[3] == ""{
				continue
			}
			processName := strings.Trim(tmp[0], " ") 	// 进程名
			//processName = string([]rune(processName)[1:len(processName)-1])
			processName = strings.Trim(processName, "\"")
			pid := strings.Trim(tmp[1], " ")			// pid
			pid = strings.Trim(pid, "\"")
			status := strings.Trim(tmp[5], " ")			// 状态 Unknown/Running/Not responding 或许还有
			status = strings.Trim(status, "\"")
			windowName := strings.Trim(tmp[8], " ")		// 窗口名称
			windowName = strings.Trim(windowName, "\"")

			//截取出来的数组中的每个元素都带有双引号，所以做字符串截取处理
			tmp[0] = processName
			tmp[1] = pid
			tmp[2] = status
			tmp[3] = windowName

			tasklist = append(tasklist, tmp)
		}
		counter ++
	}
	return tasklist
}

//根据名字返回有多少个同名进程，一个浏览器多个tab都等于多个进程；java、tomcat等的貌似都叫java.exe
//func CountWinTask(taskName string, taskList [][]string) int  {
//
//}
