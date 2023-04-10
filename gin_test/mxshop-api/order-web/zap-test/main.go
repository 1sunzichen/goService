package main

import (
	"go.uber.org/zap"
)

func main(){
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	url:="www.baidu.com"
	logger.Info("failer to fetch Url",
		zap.String("url",url),
		zap.Int("attempt",3),
		)
	//sugar := logger.Sugar()

	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
}
