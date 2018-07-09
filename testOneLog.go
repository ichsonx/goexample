package main

import (
	"github.com/francoispqt/onelog"
	"os"
	"time"
	"errors"
)

var logger *onelog.Logger

func main() {
	helloLevels()
	//helloHook()
	//helloExtraFields()
}

func helloBasic()  {
	// create a new Logger
	// first argument is an io.Writer
	// second argument is the level, which is an integer
	logger := onelog.New(
		os.Stdout,
		onelog.ALL, // shortcut for onelog.DEBUG|onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL,
	)
	logger.Info("hello world !") // {"level":"info","message":"hello world"}
}

/*
	onelog允许创建实体的时候设置日志的等级，如果设置为高级，则即使在代码中使用了更低级（例如debug）的日志方法，信息都不会被输出出来的
    onelog.DEBUG
    onelog.INFO
    onelog.WARN
    onelog.ERROR
    onelog.FATAL
	这里是按照你所设置的等级来打印信息的。即使你包括了最低debug等级，但没包含warn，也不会输出warn信息

*/
func helloLevels()  {
	logger = onelog.New(
		os.Stdout,
		onelog.INFO|onelog.WARN|onelog.FATAL,
	)
	//日志级别低于声明的界别，则会抛弃
	logger.Debug("helloLevels")
	logger.Fatal("helloLevels")
	logger.Warn("helloLevels")
	logger.Error("helloLevels")
}

/*
	设置的hook会应用到所有的日志信息。这里time不能定义到最前面比较不人性化。
*/
func helloHook()  {
	logger := onelog.New(
		os.Stdout,
		onelog.ALL,
	)
	logger.Hook(func(e onelog.Entry) {
		e.String("time", time.Now().Format(time.RFC3339))
	})
	logger.Info("hello world !") // {"level":"info","message":"hello world","time":"2018-05-06T02:21:01+08:00"}
}

/*
	可以额添加一些字段甚至对象（对象依然以json内嵌对象格式展示）,不过需要给额外添加的字段静态声明是什么类型
*/
func helloExtraFields()  {
	logger := onelog.New(
		os.Stdout,
		onelog.ALL,
	)

	logger.DebugWithFields("i'm not sure what's going on", func(e onelog.Entry) {
		e.String("string", "foobar")
		e.Int("int", 12345)
		e.Int64("int64", 12345)
		e.Float("float64", 0.15)
		e.Bool("bool", true)
		e.Err("err", errors.New("someError"))
		e.ObjectFunc("user", func(e onelog.Entry) {
			e.String("name", "somename")
		})
	})
}