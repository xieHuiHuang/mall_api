/**
 * @file: main.go
 * @time: 2022-10-13 16:50
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package main

import (
	"fmt"
	"mall_api/user_web/global"

	"go.uber.org/zap"

	"mall_api/user_web/initialize"
)

func main() {
	//Port := 8021

	//1. 初始化logger
	initialize.InitLogger()

	//2. 初始化配置文件
	initialize.InitConfig()

	//3. 初始化routers
	Router := initialize.Routers()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//5. 初始化srv的连接
	//initialize.InitSrvConn()

	//viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	//debug := viper.GetBool("MALL_DEBUG")
	//if !debug {
	//	port, err := utils.GetFreePort()
	//	if err == nil {
	//		global.ServerConfig.Port = port
	//	}
	//}

	//注册验证器
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
	//	_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
	//		return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
	//	}, func(ut ut.Translator, fe validator.FieldError) string {
	//		t, _ := ut.T("mobile", fe.Field())
	//		return t
	//	})
	//}

	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	//zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
