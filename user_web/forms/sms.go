/**
 * @file: sms.go
 * @time: 2022-10-13 17:31
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`  // 1代表注册、2代表动态验证登录
	//1. 注册发送短信验证码和动态验证码登录发送验证码
}
