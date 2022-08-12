package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserID string
	jwt.RegisteredClaims
}

type JwtUtil struct {
	SecretKey string
	ExpiresAt time.Duration
}

func New(secretKey string, expiresAt time.Duration) *JwtUtil {
	return &JwtUtil{
		secretKey,
		expiresAt,
	}
}

func (j *JwtUtil) GenerateToken(userID string) (*string, error) {
	return nil, nil
}

func (j *JwtUtil) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
