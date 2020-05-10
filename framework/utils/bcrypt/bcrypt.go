package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparedPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}
