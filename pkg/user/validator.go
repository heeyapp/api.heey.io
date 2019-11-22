package user

import "regexp"

// ValidateEmail 验证邮箱
func ValidateEmail(email string) bool {
	return regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`).MatchString(email)
}

// ValidatePassWord 验证密码
func ValidatePassWord(password string) bool {
	return regexp.MustCompile(`^\S{8,}$`).MatchString(password)
}
