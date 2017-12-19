package main

import (
	"path/filepath"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {
	/*
		至少于目前来说，使用相对路径读取配置文件是没问题的。
		网上还提到了：os.Executable() // 获得程序路径、filepath.Abs()等拼接路径的方法
	*/
	file, err := os.Open("./config/targetlist.json")
	check(err)
	fmt.Println(file.Name())
	//下面这个方法同样可以获取路径，不过是执行文件的绝对路径
	fmt.Println(getCurrentDirectory())
}
func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

