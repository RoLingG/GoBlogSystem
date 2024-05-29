package log_stash_v2

import "github.com/dgrijalva/jwt-go/v4"

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	RoleID   int    `json:"role"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

// 这里解析token的方法是跳过了密钥认证，也就是说无论解析是否成功/失败，都会获取token内携带用户信息/为空信息，但不安全
// 而且这个方法只返回token的自定义荷载内容(也就是用户信息)，而不是ParseToken那样获取token内的自定义荷载内容和标准声明
// 且这个方法不在乎token的失效时间这些标准声明内容，不用去验证。而ParseToken需要验证里面是否有哪些内容
func parseToken(token string) (jwtPayLoad *JwtPayLoad) {
	//解析传过来的token，获取解析后的token。&CustomClaims{}用于存放解析后的JWT载荷（payload）和标准声明（claims）
	//func(token *jwt.Token) (interface{}, error) { return []byte(""), nil }验证JWT的签名
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	//如果token不合法，则token携带的内容判定为空
	if Token == nil || Token.Claims == nil {
		return nil
	}
	//token不为空，则获取解析后的token的声明，并将其断言成我们自己设置的CustomClaims结构体
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	//返回声明中的JwtPayLoad结构体内的值，也就是token内携带的用户信息
	return &claims.JwtPayLoad
}
