package main

import (
	"github.com/sirupsen/logrus"
	"os"
)
/*
地址：https://github.com/Sirupsen/logrus
一个简书例子：https://www.jianshu.com/p/5fac8bed4505
一个stars很多的项目。日志项目。
comment：hook暂时来说作用不大。如果要一个日志要同时输出到多个地方（console、logfile）只能创建多个实例来完成。
*/
func main() {
	useLogToMutiplePlace()
}

//可以将日志输出到文件中，但不会出现在console中
func useLogToMutiplePlace()  {
	log := logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败")


}
