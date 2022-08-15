package hash

import (
	BcryptHelper "crosscheck-golang/app/utils/bcrypt"
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
	return nil, nil
}

func (h *HashImpl) ComparePassword(hashPassword string, password string) error {
	return nil
}
