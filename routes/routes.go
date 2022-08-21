package routes

import (
	authRepository "crosscheck-golang/app/features/authentication/data/repository"
	authpersistent "crosscheck-golang/app/features/authentication/data/source/persistent"
	authloginuc "crosscheck-golang/app/features/authentication/domain/usecase/login"
	authregistrationuc "crosscheck-golang/app/features/authentication/domain/usecase/registration"
	authcontroller "crosscheck-golang/app/features/authentication/presentation/http/controller"
	authrouter "crosscheck-golang/app/features/authentication/presentation/http/router"
	"crosscheck-golang/app/utils/bcrypt"
	"crosscheck-golang/app/utils/clock"
	"crosscheck-golang/app/utils/hash"
	jwtUtils "crosscheck-golang/app/utils/jwt"
	"crosscheck-golang/config"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Route struct {
	app    *echo.Echo
	db     *sqlx.DB
	config *config.Config
}

func New(app *echo.Echo, db *sqlx.DB, config *config.Config) *Route {
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
	accessToken := jwtUtils.New(r.config.AccessToken.Secret, r.config.AccessToken.Expires)
	refreshToken := jwtUtils.New(r.config.RefreshToken.Secret, r.config.RefreshToken.Expires)
	hash := hash.New(bcrypt.New())
	clock := clock.New()
	authPersistent := authpersistent.New(r.db)
	authRepository := authRepository.New(authPersistent, clock)
	authRegistrationUsecase := authregistrationuc.New(authRepository, accessToken, refreshToken, hash)
	authLoginUsecase := authloginuc.New(authRepository, accessToken, refreshToken, hash)
	authController := authcontroller.New(authRegistrationUsecase, authLoginUsecase)
	authRouter := authrouter.New(r.app, authController)
	authRouter.Run()
}
