package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

// 返回结构体中msg的参数
func GetValidMsg(err error, obj any) string {
	//获取obj的指针
	getobj := reflect.TypeOf(obj)
	//断言err接口为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每个错误信息
			//根据报错字段名去获取结构体的相关字段
			if f, exits := getobj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

//检验url是否合法
//func ValidateURL(str string) (bool, error) {
//	u, err := url.ParseRequestURI(str)
//	if err != nil {
//		return false, err
//	}
//	if u.Scheme == "" || u.Host == "" {
//		return false, fmt.Errorf("invalid URL format: scheme and host are required")
//	}
//	return true, nil
//}

// 正则检验url是否合法
func ValidateURL(str string) bool {
	// 正则表达式用于匹配一个简单的 URL 格式
	// 这个正则表达式并不涵盖所有可能的 URL 格式，但可以作为基本的验证
	re := regexp.MustCompile(
		`^(https?|ftp):\/\/[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]$`,
	)
	return re.MatchString(str)
}
