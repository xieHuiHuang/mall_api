/**
 * @file: base.go
 * @time: 2022-10-13 17:02
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package router

import (
	"github.com/gin-gonic/gin"
	"mall_api/user_web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}

}
