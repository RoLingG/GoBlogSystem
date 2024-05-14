package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd 加密密码 输入时密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// CheckPwd 验证密码 hash加密之后的密码  输入的密码
func CheckPwd(hashedPwd string, pwd string) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
