package authrouter

import (
	authcontroller "crosscheck-golang/app/features/authentication/presentation/http/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	echo       *echo.Echo
	controller *authcontroller.AuthController
}

func New(echo *echo.Echo, controller *authcontroller.AuthController) *AuthRouter {
	return &AuthRouter{
		echo,
		controller,
	}
}

func (s *AuthRouter) Run() {
	router := s.echo.Group("/auth")
	// swagger:route POST /auth/registration authentication registration
	//
	// Register a new user.
	//
	// consumes:
	// 	- application/x-www-form-urlencoded
	// produces:
	// 	- application/json
	//
	// responses:
	// 	200: RegistrationSuccessResponse
	//  500: Registration500Response
	//  400: Registration400Response
	router.POST("/registration", s.controller.Registration)
	router.POST("/login", s.controller.Login)
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Auth")
	})
}
