package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server       ServerConfig
	DbConfig     DBConfig
	AccessToken  TokenConfig
	RefreshToken TokenConfig
}

type ServerConfig struct {
	AppVersion   string
	Port         string
	Env          string
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

type TokenConfig struct {
	Name     string
	Path     string
	Secret   string
	Secure   bool
	HttpOnly bool
	Expires  time.Duration
}

func New(filename ...string) *Config {
	if err := godotenv.Load(filename...); err != nil {
		log.Println("Failed to load .env file")
		log.Fatal(err)

		return nil
	}

	dbConfig := GetDBConfig()

	serverConfig, err := GetServerConfig()
	if err != nil {
		log.Println("Failed to load server config")
		log.Fatal(err)

		return nil
	}

	tokenCookieConfig, err := GetTokenCookieConfig()
	if err != nil {
		log.Println("Failed to load token cookie config")
		log.Fatal(err)

		return nil
	}

	refreshTokenCookieConfig, err := GetRefreshTokenCookieConfig()
	if err != nil {
		log.Println("Failed to load refresh token cookie config")
		log.Fatal(err)

		return nil
	}

	return &Config{
		*serverConfig,
		*dbConfig,
		*tokenCookieConfig,
		*refreshTokenCookieConfig,
	}
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

func GetTokenCookieConfig() (*TokenConfig, error) {
	secure, err := strconv.ParseBool(os.Getenv("JWT_ACCESS_TOKEN_SECURE"))
	if err != nil {
		return nil, err
	}

	httpOnly, err := strconv.ParseBool(os.Getenv("JWT_ACCESS_TOKEN_HTTPONLY"))
	if err != nil {
		return nil, err
	}

	expires, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRES"))
	if err != nil {
		return nil, err
	}

	return &TokenConfig{
		Name:     os.Getenv("JWT_ACCESS_TOKEN_NAME"),
		Path:     os.Getenv("JWT_ACCESS_TOKEN_PATH"),
		Secret:   os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
		Secure:   secure,
		HttpOnly: httpOnly,
		Expires:  time.Duration(expires) * time.Millisecond,
	}, nil
}

func GetRefreshTokenCookieConfig() (*TokenConfig, error) {
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

	return &TokenConfig{
		Name:     os.Getenv("JWT_REFRESH_TOKEN_NAME"),
		Path:     os.Getenv("JWT_REFRESH_TOKEN_PATH"),
		Secret:   os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		Secure:   secure,
		HttpOnly: httpOnly,
		Expires:  time.Duration(expires) * time.Millisecond,
	}, nil
}
