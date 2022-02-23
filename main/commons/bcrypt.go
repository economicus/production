package commons

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const cost = 12

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func ComparePassword(nonHashed, hashed []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, nonHashed)
	if err != nil {
		return fmt.Errorf("wrong password")
	}
	return nil
}
