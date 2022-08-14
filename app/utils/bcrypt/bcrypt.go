package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {
	pw := []byte(password)
	encrypted, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result := string(encrypted)
	return &result, nil
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
