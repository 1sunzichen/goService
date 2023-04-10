package main

import (
	"go.uber.org/zap"
)

func NewLogger()(*zap.Logger,error){
	cfg:=zap.NewProductionConfig()
	cfg.OutputPaths=[]string{
		"./pro.log",
	}
	return cfg.Build()
}
func main(){
	logger, _ := NewLogger()
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
