package utils

import (
	"fmt"
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
	if err := j.validateStructProperty(); err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, fmt.Errorf("UserID is required")
	}

	claims := &JwtClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(j.ExpiresAt),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (j *JwtUtil) ValidateToken(token string) (*jwt.Token, error) {

	if err := j.validateStructProperty(); err != nil {
		return nil, err
	}

	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})
}

// Validate property of JwtUtil struct privately
func (j *JwtUtil) validateStructProperty() error {
	if j.SecretKey == "" {
		return fmt.Errorf("SecretKey is required")
	}

	if j.ExpiresAt == 0 {
		return fmt.Errorf("ExpiresAt must be greater than 0")
	}

	return nil
}
