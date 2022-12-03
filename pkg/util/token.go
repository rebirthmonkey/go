package util

import "golang.org/x/crypto/bcrypt"

func GenerateToken() string {
	return RandomString(24)
}

func EncryptPassword(password string) (string, error) {
	newpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(newpassword), nil
}
