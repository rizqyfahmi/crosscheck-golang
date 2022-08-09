package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := os.Getenv("APP_PORT")
	appUrl := os.Getenv("APP_URL")
	appEnv := os.Getenv("APP_ENV") // load from "APP_ENV=local go run ./console/main.go"

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	log.Printf("Open here: %s (%s)\n", appUrl, appEnv)

	if err = app.Start(":" + appPort); err != nil {
		log.Fatal("Something went wrong...")
	}

}
