package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 //QQ登录
	SignGitee SignStatus = 2 //Gitee登录
	SignEmail SignStatus = 3 //Email登录
)

// 枚举，算是gorm里面的一个重要知识
func (sign SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(sign.String())
}

func (sign SignStatus) String() string {
	var str string
	switch sign {
	case SignQQ:
		str = "QQ登录"
	case SignGitee:
		str = "Gitee登录"
	case SignEmail:
		str = "Email登录"
	default:
		str = "未知登录方式"
	}
	return str
}
