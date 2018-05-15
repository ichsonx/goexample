package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)
/*
地址：https://github.com/Sirupsen/logrus
一个简书例子：https://www.jianshu.com/p/5fac8bed4505
一个stars很多的项目。日志项目。
comment：hook暂时来说作用不大。如果要一个日志要同时输出到多个地方（console、logfile）只能创建多个实例来完成。
如无特殊情况，用法还是挺简单的。
*/
func main() {
	//useLogToMutiplePlace()
	changeFormatter()
	//useLogToMutiplePlace2()
}

//可以将日志输出到文件中，但不会出现在console中
func useLogToMutiplePlace()  {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
func useLogToMutiplePlace2()  {
	logger := log.New()
	logger.Out = os.Stdout
	logger.Formatter = &log.JSONFormatter{}
	logger.Level = log.InfoLevel

	logger.Info("just a test for log.New() method to create a logger instance")
}

//更换输出格式  json / text
//当使用text模式，time属性会是第一个字段；使用json模式，time属性非第一个字段。
func changeFormatter()  {
	//默认情况下使用的是textformatter
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&log.JSONFormatter{})
	//log := log.New()
	log.Info("just a test for json formatter")
}

