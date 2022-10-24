/**
 * @file: router.go
 * @time: 2022-10-23 12:11
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package initialize

import (
	"github.com/gin-gonic/gin"
	"mall_api/goods_web/middlewares"
	"mall_api/goods_web/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())
	//添加链路追踪
	ApiGroup := Router.Group("/goods/v1")
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)

	return Router
}
