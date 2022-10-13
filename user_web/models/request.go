/**
 * @file: request.go
 * @time: 2022-10-13 17:22
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
