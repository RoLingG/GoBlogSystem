package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中Payload的数据
type JwtPayLoad struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name,"`
	Role     int    `json:"role"` //用户权限 1.管理员 2.普通用户 3.游客
	UserID   uint   `json:"user_id"`
}

// CustomClaims jwt声明
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var MySecret []byte
