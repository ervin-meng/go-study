package main

import "go.uber.org/zap"

func main() {
	//logger,_ := zap.NewProduction() //生产环境用法
	//logger,_ := zap.NewDevelopment() //开发环境用法
	logger, _ := zapLogFile()
	//defer logger.Sync() //刷新缓存
	//name := "go-study"
	//logger.Info(name, zap.String("key", "key"), zap.Int("value", 3))
	//sugar := logger.Sugar()
	//sugar.Infow("a", "b", name)
	zap.ReplaceGlobals(logger)
	zap.S().Debug("hahah") //S()获取一个全局suger
}

func zapLogFile() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"./log/zap.log",
		"stderr",
	}
	return cfg.Build()
}
