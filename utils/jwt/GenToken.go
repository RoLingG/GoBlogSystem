package jwt

import (
	"GoRoLingG/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// GenToken 创建Token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.JWT.Secret)
	claim := CustomClaims{
		user, //jwt的payload数据
		jwt.StandardClaims{ //jwt的一些设置
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.JWT.Expires))), //默认按小时为单位进行过期，过期时间在setting.yaml的jwt内的expires设置
			Issuer:    global.Config.JWT.Issuer,                                                     //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //NewWithClaims()会生成一个jwt框架，也就是创建一个token实例
	completeToken, err := token.SignedString(MySecret)        //token对象调用SignedString(Secret)会通过Secret密钥给jwt签名，生成token字符串。这个过程先生成待签名字符串，然后生成签名，最后组合在一起形成完整的jwt
	return completeToken, err
	//生成的jwt用.对header、Payload、Signature进行拼接header.Payload.Signature
}
