package desensitization

import "regexp"

// DesensitizationTel 脱敏手机号
func DesensitizationTel(tel string) string {
	//判断手机号位数
	if len(tel) != 11 {
		return ""
	}
	re := regexp.MustCompile(`(\d{3})\d{4}(\d{4})`)
	telRe := re.ReplaceAllString(tel, "$1****$2")
	return telRe
}
