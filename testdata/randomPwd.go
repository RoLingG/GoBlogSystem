package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 生成指定长度的随机字符串
func generateRandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		randInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[randInt.Int64()]
	}
	return string(ret), nil
}

func main() {
	// 生成一个长度为10的随机字符串
	randomStr, err := generateRandomString(10)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}
	fmt.Println("Random String:", randomStr)
}
