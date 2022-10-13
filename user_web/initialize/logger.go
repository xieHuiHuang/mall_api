/**
 * @file: logger.go
 * @time: 2022-10-13 17:38
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package initialize

import "go.uber.org/zap"

func InitLogger() {
	//logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
