/**
 * @file: global.go
 * @time: 2022-10-23 12:10
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package global

import (
	ut "github.com/go-playground/universal-translator"
	"mall_api/goods_web/config"
	"mall_api/goods_web/proto"
)

var (
	Trans ut.Translator

	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	NacosConfig    *config.NacosConfig  = &config.NacosConfig{}
	GoodsSrvClient proto.GoodsClient
)
