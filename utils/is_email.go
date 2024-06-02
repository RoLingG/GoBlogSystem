package utils

import "regexp"

// IsValidEmail 检查邮箱地址是否有效
func IsValidEmail(email string) bool {
	// 邮箱的正则表达式
	emailRegex := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)
	// 使用正则表达式匹配邮箱
	return re.MatchString(email)
}
