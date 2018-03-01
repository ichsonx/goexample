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
	for _, v := range list{
		if strings.Contains(v[3], "goexample"){
			fmt.Println(v[3])
		}
	}
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
	//fmt.Printf("%s ", string(output))
	for _, line := range strings.Split(string(output), "\n"){
		//因为windows开始的2行都是些title
		if counter > 2{
			tmp := strings.Split(line, ",\"")
			var result []string
			//白痴的windows，竟然还会有空！？他妈在逗我
			if tmp[0] == "" || tmp[1] == "" || tmp[5] == "" || tmp[8] == ""{
				continue
			}
			//processName = string([]rune(processName)[1:len(processName)-1])
			//一般windows的字符串都有中文或者乱码，可以用mahonia来转码
			//以下能得到干净的值存入二维数组，但不优雅。因为windows输出的就是一堆字符串
			result = append(result, string([]rune(tmp[0])[1:len(tmp[0])-1]))	// 进程名
			result = append(result, string([]rune(tmp[1])[0:len(tmp[1])-1]))	// pid
			result = append(result, string([]rune(tmp[5])[0:len(tmp[5])-1]))	// 状态 Unknown/Running/Not responding 或许还有
			result = append(result, string([]rune(tmp[8])[:len(tmp[8])-2]))		// 窗口名称

			tasklist = append(tasklist, result)
		}
		counter ++
	}
	return tasklist
}


//根据名字返回有多少个同名进程，一个浏览器多个tab都等于多个进程；java、tomcat等的貌似都叫java.exe
//func CountWinTask(taskName string, taskList [][]string) int  {
//
//}
