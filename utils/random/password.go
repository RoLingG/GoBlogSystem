package random

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
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
