package utils

import (
	"crosscheck-golang/app/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptHashPasswordFailure(t *testing.T) {
	hashPassword := utils.HashPassword("HelloPassword")

	assert.NotNil(t, bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte("OllaPassword")))
}

func TestBcryptHashPasswordSuccess(t *testing.T) {
	hashPassword := utils.HashPassword("HelloPassword")

	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte("HelloPassword")))
}

func TestBcryptComparePasswordFailure(t *testing.T) {
	pw := []byte("HelloPassword")
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, utils.ComparePassword(string(result), "OllaPassword"))
}

func TestBcryptComparePasswordSuccess(t *testing.T) {
	pw := []byte("HelloPassword")
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, utils.ComparePassword(string(result), "HelloPassword"))
}
