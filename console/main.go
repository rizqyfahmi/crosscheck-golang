package main

import (
	"crosscheck-golang/config"
	"crosscheck-golang/config/database"
	"crosscheck-golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configuration := config.New()

	db := database.NewPostgres(configuration)
	defer db.Close()

	app := echo.New()

	routes.New(app, db, configuration).Run()

}
