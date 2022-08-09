package main

import (
	"crosscheck-golang/config"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	configuration, err := config.NewConfig()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	if err = app.Start(":" + configuration.Server.Port); err != nil {
		log.Fatal("Something went wrong...")
	}

}
