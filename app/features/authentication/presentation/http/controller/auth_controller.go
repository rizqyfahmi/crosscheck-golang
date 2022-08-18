package authcontroller

import (
	registrationusecase "crosscheck-golang/app/features/authentication/domain/usecase/registration"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	registraion registrationusecase.RegistrationUsecase
}

func New(registraion registrationusecase.RegistrationUsecase) *AuthController {
	return &AuthController{
		registraion: registraion,
	}
}

func (controller *AuthController) Registration(c echo.Context) error {
	return nil
}
