package main

import (
	"time"
	"go.uber.org/zap"
)

var url = "www.163.com"

func main() {
	useSugaredLogger()
	uselogger()
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
