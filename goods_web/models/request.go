/**
 * @file: request.go
 * @time: 2022-10-23 12:46
 * @Author: jack
 * @Email: 793936517@qq.com
 * @desc:
 **/

package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
