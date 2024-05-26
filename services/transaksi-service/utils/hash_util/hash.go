package hash_util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPin(pinA string, pinB string) error {
	return bcrypt.CompareHashAndPassword([]byte(pinA), []byte(pinB))
}
