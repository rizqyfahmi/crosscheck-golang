package utils

import (
	"crosscheck-golang/app/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenEmptySecretKey(t *testing.T) {
	jwt := utils.New("", 10*time.Minute)
	result, err := jwt.GenerateToken("UserID")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "SecretKey is required")
}

func TestGenerateTokenEmptyExpiresAt(t *testing.T) {
	jwt := utils.New("SecretKey", 0)
	result, err := jwt.GenerateToken("UserID")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "ExpiresAt must be greater than 0")
}

func TestGenerateTokenEmptyParameter(t *testing.T) {
	jwt := utils.New("SecretKey", 10*time.Minute)
	result, err := jwt.GenerateToken("")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "UserID is required")
}

func TestGenerateTokenSuccess(t *testing.T) {
	jwt := utils.New("SecretKey", 10*time.Minute)
	result, err := jwt.GenerateToken("UserID")

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestValidateTokenEmptySecretKey(t *testing.T) {
	jwt := utils.New("", 30*time.Minute)

	token, err := jwt.ValidateToken("Token")

	assert.Nil(t, token)
	assert.Error(t, err)
	assert.EqualError(t, err, "SecretKey is required")
}

func TestValidateTokenEmptyExpiresAt(t *testing.T) {
	jwt := utils.New("SecretKey", 0*time.Minute)

	token, err := jwt.ValidateToken("Token")

	assert.Nil(t, token)
	assert.Error(t, err)
	assert.EqualError(t, err, "ExpiresAt must be greater than 0")
}

func TestValidateTokenEmptyParameter(t *testing.T) {
	jwt := utils.New("SecretKey", 30*time.Minute)

	token, err := jwt.ValidateToken("")

	assert.Nil(t, token)
	assert.Error(t, err)
	assert.EqualError(t, err, "token is required")
}

func TestValidateTokenInvalid(t *testing.T) {
	jwt := utils.New("SecretKey", 30*time.Minute)

	token, err := jwt.ValidateToken("Token")

	assert.Nil(t, token)
	assert.Error(t, err)
	assert.EqualError(t, err, "token contains an invalid number of segments")
}

func TestValidateTokenSuccess(t *testing.T) {
	jwt := utils.New("SecretKey", 30*time.Minute)
	tokenString, err := jwt.GenerateToken("UserID")
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.ValidateToken(*tokenString)

	assert.NotNil(t, token)
	assert.Nil(t, err)
}
