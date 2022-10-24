/**
 * @file: banner.go
 * @time: 2022-10-23 12:10
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package forms

type BannerForm struct {
	Image string `form:"image" json:"image" binding:"url"`
	Index int    `form:"index" json:"index" binding:"required"`
	Url   string `form:"url" json:"url" binding:"url"`
}
