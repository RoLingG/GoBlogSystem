package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (arr *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*arr = []string{}
		return nil
	}
	*arr = strings.Split(string(v), "\n")
	return nil
}

// Value 存入数据库
func (arr Array) Value() (driver.Value, error) {
	//将数字转换成值
	return strings.Join(arr, "\n"), nil
}
