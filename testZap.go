package main

import (
	"time"
	"go.uber.org/zap"
	"fmt"
	"encoding/json"
	"go.uber.org/zap/zapcore"
	"log"
)

var url = "www.163.com"
var Logger *zap.Logger

/*
uber 的开源日志项目。
地址：https://github.com/uber-go/zap
*/
func main() {
	//useSugaredLogger()
	//uselogger()
	//useLogToFile()
	//useSimpleLogger()
	useSimpleSugar()
}

//自定义的config，输出到日志文件。需要修改zap的config，然后从config.buile一个logger出来。
func useSimpleLogger()  {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"zap-simple.log"}
	logger, _ := cfg.Build()
	defer logger.Sync()

	logger.Info("test simple log", zap.String("somekey", "value"))
}

//自定义的config，输出到日志文件。需要修改zap的config，然后从config.buile一个logger出来。
//在从logger生成一个sugar来使用，在log中插入键-值，比单纯使用logger更方便一点。
func useSimpleSugar()  {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"zap-simple.log"}
	logger, _ := cfg.Build()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("test simple log",
		"somekey", "value",
			"url", url,
	)
}



/*
In contexts where performance is nice, but not critical, use the SugaredLogger.
It's 4-10x faster than than other structured logging packages and includes both structured and printf-style APIs.
当对性能要求不是很严格的时候，可以用SugaredLogger
*/
func useSugaredLogger()  {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

/*
When performance and type safety are critical, use the Logger.
It's even faster than the SugaredLogger and allocates far less, but it only supports structured logging.
当对性能有严格要求的时候就使用Logger。比SugaredLogger更快，但就只支持一种结构的日志。
*/
func uselogger()  {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

//只能输出到file，不能同时输出console、file
func useLogToFile()  {
	initLogger("out.log", "DEBUG", false)
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	Logger.Info("info:", zap.String("s", s[0]))
	Logger.Error("info:", zap.String("s", s[1]))
	Logger.Debug("info:", zap.String("s", s[2]))
	Logger.Fatal("info:", zap.String("s", s[3]))
}
func initLogger(lp string, lv string, isDebug bool) {
	var js string
	if isDebug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, lv)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, lv, lp, lp)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}


