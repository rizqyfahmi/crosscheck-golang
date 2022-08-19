package authcontroller

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	registrationusecase "crosscheck-golang/app/features/authentication/domain/usecase/registration"
	"log"
	"net/http"

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
	param := new(param.RegistrationParam)

	log.Printf("\nRegistration controller: binding request parameters")
	if err := c.Bind(param); err != nil {
		log.Printf("\nRegistration controller: error binding! -> %+v", err)
		return c.JSON(http.StatusInternalServerError, exception.InternalServerError)
	}

	log.Printf("\nRegistration controller: validating request parameters")
	if err := c.Validate(param); err != nil {
		log.Printf("\nRegistration controller: error validation! -> %+v", err)
		return c.JSON(http.StatusBadRequest, exception.BadRequest)
	}

	log.Printf("\nRegistration controller: executing registration usecase")
	authEntity, err := controller.registraion.Call(*param)

	if err != nil {
		log.Printf("\nRegistration controller: error usecase! -> %+v", err)
		return c.JSON(http.StatusBadRequest, exception.BadRequest)
	}

	log.Println("Completed!")
	return c.JSON(http.StatusOK, authEntity)

}
