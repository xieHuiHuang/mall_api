/**
 * @file: config.go
 * @time: 2022-10-13 17:38
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
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
	//从配置文件中读取出对应的配置
	debug := GetEnvInfo("MALL_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user_web/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：&v", global.NacosConfig)
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,                         //连接超时时间
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",   //日志目录
		CacheDir:            "tmp/nacos/cache", //nacos缓存目录
		LogLevel:            "debug",           //日志级别
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败：%s", err.Error())
	}
	fmt.Println(&global.ServerConfig)
}

func InitConfig2() {
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
