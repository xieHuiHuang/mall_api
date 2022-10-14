/**
 * @file: router.go
 * @time: 2022-10-13 17:06
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package initialize

import (
	"github.com/gin-gonic/gin"
	"mall_api/user_web/middlewares"
	"mall_api/user_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//Router.GET("/health", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    http.StatusOK,
	//		"success": true,
	//	})
	//})

	//配置跨域
	Router.Use(middlewares.Cors())
	
	ApiGroup := Router.Group("/user/v1")
	router.InitUserRouter(ApiGroup)
	//router.InitBaseRouter(ApiGroup)

	return Router
}
