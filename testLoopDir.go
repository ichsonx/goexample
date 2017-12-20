package main

import (
	"path/filepath"
	"os"
	"strings"
	"fmt"
	"io/ioutil"
)

/*
	遍历目录下的子目录及文件
filepath.walk：树形遍历全部的子目录和文件，直到尽头。
ioutil.ReadDir：列出该目录下所有子文件夹和文件，只一层，不再往下。
*/

func main() {
	useReadDir()
	//useWalk()
}

//walk方法会递归遍历，直到再没有子目录为止。他会使用一个walkfunc方法来处理遍历的所有路径。
//walkfunc方法的参数和返回值必须是func(string, os.FileInfo, error)(error)，这种模式
func useWalk()  {
	root := "e:\\pycharmWorkspace"
	filepath.Walk(root, walkfunc)
}
func walkfunc(path string, f os.FileInfo, err error)  error {
	if strings.Contains(f.Name(), "spider") && !f.IsDir(){
		fmt.Println(path)
	}
	return nil
}

//此方法是用ioutil.ReadDir来读取指定目录下的所有子文件夹和文件（只一层，不会再深入）。获得os.FileInfo列表。如果要递归，需要自己补充写代码。
func useReadDir()  {
	root := "e:\\pycharmWorkspace"
	files, _ := ioutil.ReadDir(root)
	for _, file := range files{
		fmt.Println(file.Name())
	}
}
