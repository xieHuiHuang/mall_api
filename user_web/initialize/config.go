/**
 * @file: config.go
 * @time: 2022-10-13 17:38
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall_api/user_web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}

func InitConfig() {
	debug := GetEnvInfo("MALL_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user_web/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//serverConfig := config.ServerConfig{}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	fmt.Println(global.ServerConfig)
	zap.S().Infof("配置信息：&v", global.ServerConfig)
	fmt.Printf("%v", v.Get("name"))
	//viper的功能 -动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Println("config file channed: ", e.Name)
		zap.S().Infof("配置文件产生变化： %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		//fmt.Println(global.ServerConfig)
		zap.S().Infof("配置信息： &v", global.ServerConfig)
	})
}
