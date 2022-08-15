package hash

import (
	BcryptHelper "crosscheck-golang/app/utils/bcrypt"

	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	HashPassword(password string) (*string, error)
	ComparePassword(hashPassword string, password string) error
}

type HashImpl struct {
	hash BcryptHelper.Bcrypt
}

func New(hash BcryptHelper.Bcrypt) Hash {
	return &HashImpl{
		hash,
	}
}

func (h *HashImpl) HashPassword(password string) (*string, error) {
	pw := []byte(password)
	encrypted, err := h.hash.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result := string(encrypted)
	return &result, nil
}

func (h *HashImpl) ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := h.hash.CompareHashAndPassword(hw, pw)
	return err
}
