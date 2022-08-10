package routes

import (
	"crosscheck-golang/config"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	App    *echo.Echo
	Db     *sql.DB
	Config *config.Config
}

func New(app *echo.Echo, db *sql.DB, config *config.Config) *Route {
	return &Route{
		app,
		db,
		config,
	}
}

func (r *Route) Run() {

	r.getGeneralRoute()
	r.getAuthRoute()

	if err := r.App.Start(":" + r.Config.Server.Port); err != nil {
		log.Fatal("Something went wrong...")
	}
}

// Get general route privately
func (r *Route) getGeneralRoute() {
	router := r.App
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}

// Get auth route privately
func (r *Route) getAuthRoute() {
	router := r.App.Group("/auth")

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Auth")
	})
}
