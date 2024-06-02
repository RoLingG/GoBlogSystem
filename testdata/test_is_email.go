package main

import (
	"fmt"
	"regexp"
)

// 检查邮箱地址是否有效
func isValidEmail(email string) bool {
	// 邮箱的正则表达式
	emailRegex := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)
	// 使用正则表达式匹配邮箱
	return re.MatchString(email)
}

func main() {
	// 测试邮箱
	testEmail := "123456@123456.com"

	// 检测邮箱是否有效
	if isValidEmail(testEmail) {
		fmt.Println(testEmail, "是有效的邮箱地址")
	} else {
		fmt.Println(testEmail, "不是有效的邮箱地址")
	}
}
