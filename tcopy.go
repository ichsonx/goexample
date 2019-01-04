package main

import (
	"log"
	"os"
)

func main() {
	dest := "./production"
	//src := "../goexample"

	if _, err := os.Stat(dest); err == nil{
		if err := os.RemoveAll(dest); err != nil{
			log.Fatalf("\n %s 目录删除失败，退出打包...错误信息： %v\n", err)
		}
	}
	if err := os.Mkdir(dest, 0777); err != nil {
		log.Fatalf("\n %s 目录创建失败，退出打包...错误信息： %v\n", err)
	}

}