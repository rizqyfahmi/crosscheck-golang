package routes

import (
	AuthRepository "crosscheck-golang/app/features/authentication/data/repository"
	AuthLocalDataSource "crosscheck-golang/app/features/authentication/data/source"
	AuthUsecase "crosscheck-golang/app/features/authentication/domain/usecase"
	AuthController "crosscheck-golang/app/features/authentication/presentation/controller"
	"crosscheck-golang/config"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Route struct {
	app    *echo.Echo
	db     *gorm.DB
	config *config.Config
}

func New(app *echo.Echo, db *gorm.DB, config *config.Config) *Route {
	return &Route{
		app,
		db,
		config,
	}
}

func (r *Route) Run() {

	r.getGeneralRoute()
	r.getAuthRoute()

	if err := r.app.Start(":" + r.config.Server.Port); err != nil {
		log.Fatal("Something went wrong...")
	}
}

// Get general route privately
func (r *Route) getGeneralRoute() {
	router := r.app
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}

// Get auth route privately
func (r *Route) getAuthRoute() {
	localSource := AuthLocalDataSource.New(r.db)
	repository := AuthRepository.New(localSource)
	usecase := AuthUsecase.New(repository, *r.config)
	c := AuthController.New(usecase)

	router := r.app.Group("/auth")

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Auth")
	})

	router.POST("/registration", c.Registration)

}
