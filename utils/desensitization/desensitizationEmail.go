package desensitization

import "regexp"

// 脱敏邮箱
func DesensitizationEmail(email string) string {
	re := regexp.MustCompile("(^\\w)[^@]*(@.*$)")
	// 使用 ReplaceAllString 进行替换
	emailRe := re.ReplaceAllString(email, "$1****$2")
	return emailRe
}
