package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
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
