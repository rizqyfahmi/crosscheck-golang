package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// this file is used for being a helper to create selector of golang.org/x/crypto/bcrypt where it isn't provided as default
type Bcrypt interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
}

type BcryptImpl struct{}

func New() Bcrypt {
	return &BcryptImpl{}
}

func (s *BcryptImpl) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (s *BcryptImpl) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
