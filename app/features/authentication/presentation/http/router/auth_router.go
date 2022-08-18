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
	router.POST("/registration", s.controller.Registration)
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Auth")
	})
}
