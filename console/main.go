package main

import (
	validatorutil "crosscheck-golang/app/utils/validator"
	"crosscheck-golang/config"
	"crosscheck-golang/config/database"
	"crosscheck-golang/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	configuration := config.New()

	db := database.NewPostgres(configuration)
	defer db.Close()

	app := echo.New()
	app.Validator = validatorutil.New(validator.New())

	routes.New(app, db, configuration).Run()

}
