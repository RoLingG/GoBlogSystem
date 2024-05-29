package jwt

import (
	"GoRoLingG/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken  解析Token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.JWT.Secret)                                                                //将密钥byte化用于解析token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) { //这里被赋值的token已经是进过解析之后的了，通过密钥解析，更安全
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { //进行断言
		return claims, err
	}
	return nil, errors.New("invalid token")
	//从上面可以看出有三层校验，这三层分别就是对应header、payload、signature
}
