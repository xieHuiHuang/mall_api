/**
 * @file: user.go
 * @time: 2022-10-13 17:03
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package router

import (
	"github.com/gin-gonic/gin"
	"mall_api/user_web/api"
	"mall_api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("login", api.Login)
		//UserRouter.POST("register", api.Register)
		//
		//UserRouter.GET("detail", api.GetUserDetail)
		//UserRouter.PATCH("update", api.UpdateUser)
	}
	//服务注册和发现
}
