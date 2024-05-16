package random

import (
	"math/rand"
	"strings"
	"time"
)

// 验证码包含的字符集
var stringCode = "adcdefghijklmnopqrstuvwxyz0123456789"

// RandCode 随机验证码生成
func RandCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	// 创建一个空的字符串，用于存放最终的随机字符串
	var code strings.Builder
	// 生成指定长度的随机字符串
	for i := 0; i < length; i++ {
		// 生成一个随机索引，用于从字符集中选择字符
		index := rand.Intn(len(stringCode))
		// 将选中的字符添加到字符串中
		code.WriteByte(stringCode[index])
	}
	return code.String()
}
