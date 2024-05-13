package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/utils/jwt"
	"fmt"
)

func main() {
	//读取配置文件，main中调用InitConfig
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLogger()

	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserName: "test",
		NickName: "testest",
		Role:     2,
		UserID:   1,
	})
	fmt.Println(token, err)
	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ0ZXN0Iiwibmlja19uYW1lIjoidGVzdGVzdCIsInJvbGUiOjIsInVzZXJfaWQiOjEsImV4cCI6MTcxNTYxNDkzMy44NjMzMjQ5LCJpc3MiOiIxMjM0In0.o3CkZye8_gKmc4N71_U5Wzn528IKg0EiG2FPKkh7mZ8")
	fmt.Println(claims, err)
}
