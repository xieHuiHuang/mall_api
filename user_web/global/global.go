/**
 * @file: global.go
 * @time: 2022-10-13 17:05
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package global

import (
	ut "github.com/go-playground/universal-translator"
	"mall_api/user_web/config"
	//"mall_api/user_web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	//NacosConfig *config.NacosConfig = &config.NacosConfig{}

	//UserSrvClient proto.UserClient
)
