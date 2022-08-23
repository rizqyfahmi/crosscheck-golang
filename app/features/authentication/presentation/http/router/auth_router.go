package authrouter

import (
	authcontroller "crosscheck-golang/app/features/authentication/presentation/http/controller"

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
	// 	- application/json
	// produces:
	// 	- application/json
	//
	// responses:
	// 	200: AuthSuccessResponse
	//  500: InternalServerError
	//  400: BadRequest
	router.POST("/registration", s.controller.Registration)
	// swagger:route POST /auth/login authentication login
	//
	// Enter the system.
	//
	// consumes:
	// 	- application/json
	// produces:
	// 	- application/json
	//
	// responses:
	// 	200: AuthSuccessResponse
	//  500: InternalServerError
	//  400: BadRequest
	router.POST("/login", s.controller.Login)
}
