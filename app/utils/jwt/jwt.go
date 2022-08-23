package jwt

import (
	"crosscheck-golang/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserID string
	jwt.RegisteredClaims
}

type JwtUtil interface {
	GenerateToken(userID string) (*string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type AccessToken JwtUtil
type RefreshToken JwtUtil

type JwtUtilImpl struct {
	secretKey string
	expiresAt time.Duration
}

func New[T any](param config.TokenConfig) T {
	var util interface{} = &JwtUtilImpl{
		param.Secret,
		param.Expires,
	}
	return util.(T)
}

func (j *JwtUtilImpl) GenerateToken(userID string) (*string, error) {
	if err := j.validateStructProperty(); err != nil {
		return nil, err
	}

	if userID == "" {
		return nil, fmt.Errorf("UserID is required")
	}

	claims := &JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(j.expiresAt),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (j *JwtUtilImpl) ValidateToken(token string) (*jwt.Token, error) {

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
		return []byte(j.secretKey), nil
	})
}

// Validate property of JwtUtilImpl struct privately
func (j *JwtUtilImpl) validateStructProperty() error {
	if j.secretKey == "" {
		return fmt.Errorf("SecretKey is required")
	}

	if j.expiresAt == 0 {
		return fmt.Errorf("ExpiresAt must be greater than 0")
	}

	return nil
}
