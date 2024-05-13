package pwd

import (
	"fmt"
	"regexp"
)

// 正则检测密码强度
func PasswordLevel(password string) {
	// 检查是否包含大写字母、小写字母、数字和特殊字符
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`\W`).MatchString(password)

	if !hasUpper && !hasLower && !hasDigit && !hasSpecial {
		fmt.Printf("password level is weak %s\n", password)
	} else if !hasUpper && !hasLower && (hasDigit || hasSpecial) ||
		!hasUpper && hasLower && (hasDigit || hasSpecial) ||
		hasUpper && !hasLower && (hasDigit || hasSpecial) {
		fmt.Printf("password level is middle %s\n", password)
	} else {
		fmt.Printf("password level is strong %s\n", password)
	}
}
