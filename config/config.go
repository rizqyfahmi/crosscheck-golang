package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server       ServerConfig
	DbConfig     DBConfig
	Token        TokenCookieConfig
	RefreshToken RefreshTokenCookieConfig
}

type ServerConfig struct {
	AppVersion   string
	Port         string
	Env          string
	JwtSecret    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DBConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type TokenCookieConfig struct {
	Name     string
	Path     string
	Secure   bool
	HttpOnly bool
	Expires  time.Duration
}

type RefreshTokenCookieConfig struct {
	Name     string
	Path     string
	Secure   bool
	HttpOnly bool
	Expires  time.Duration
}

func NewConfig(filename ...string) (*Config, error) {
	if err := godotenv.Load(filename...); err != nil {
		return nil, err
	}

	dbConfig := GetDBConfig()

	serverConfig, err := GetServerConfig()
	if err != nil {
		return nil, err
	}

	tokenCookieConfig, err := GetTokenCookieConfig()
	if err != nil {
		return nil, err
	}

	refreshTokenCookieConfig, err := GetRefreshTokenCookieConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		*serverConfig,
		*dbConfig,
		*tokenCookieConfig,
		*refreshTokenCookieConfig,
	}, nil
}

func GetServerConfig() (*ServerConfig, error) {
	readTimeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	writeTimeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		AppVersion:   os.Getenv("APP_VERSION"),
		Port:         os.Getenv("APP_PORT"),
		Env:          os.Getenv("APP_ENV"), // Replacable with "APP_ENV=local go run ./console/main.go"
		JwtSecret:    os.Getenv("JWT_SECRET_KEY"),
		ReadTimeout:  time.Duration(readTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(writeTimeout) * time.Millisecond,
	}, nil

}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func GetTokenCookieConfig() (*TokenCookieConfig, error) {
	secure, err := strconv.ParseBool(os.Getenv("JWT_TOKEN_SECURE"))
	if err != nil {
		return nil, err
	}

	httpOnly, err := strconv.ParseBool(os.Getenv("JWT_TOKEN_HTTPONLY"))
	if err != nil {
		return nil, err
	}

	expires, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRES"))
	if err != nil {
		return nil, err
	}

	return &TokenCookieConfig{
		Name:     os.Getenv("JWT_TOKEN_NAME"),
		Path:     os.Getenv("JWT_TOKEN_PATH"),
		Secure:   secure,
		HttpOnly: httpOnly,
		Expires:  time.Duration(expires) * time.Millisecond,
	}, nil
}

func GetRefreshTokenCookieConfig() (*RefreshTokenCookieConfig, error) {
	secure, err := strconv.ParseBool(os.Getenv("JWT_REFRESH_TOKEN_SECURE"))
	if err != nil {
		return nil, err
	}

	httpOnly, err := strconv.ParseBool(os.Getenv("JWT_REFRESH_TOKEN_HTTPONLY"))
	if err != nil {
		return nil, err
	}

	expires, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRES"))
	if err != nil {
		return nil, err
	}

	return &RefreshTokenCookieConfig{
		Name:     os.Getenv("JWT_REFRESH_TOKEN_NAME"),
		Path:     os.Getenv("JWT_REFRESH_TOKEN_PATH"),
		Secure:   secure,
		HttpOnly: httpOnly,
		Expires:  time.Duration(expires) * time.Millisecond,
	}, nil
}
