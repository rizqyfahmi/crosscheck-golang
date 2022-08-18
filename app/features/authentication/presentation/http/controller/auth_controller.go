package authcontroller

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	registrationusecase "crosscheck-golang/app/features/authentication/domain/usecase/registration"
	"errors"
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
	regParam := new(param.RegistrationParam)

	if err := c.Bind(regParam); err != nil {
		log.Println("Error param!")
		log.Printf("%+v\n", regParam)
		return err
	}

	authEntity, err := controller.registraion.Call(*regParam)

	if err != nil {
		log.Println("Error usecase!")
		log.Printf("%+v\n", err)
		return errors.New(exception.BadRequest)
	}

	return c.JSON(http.StatusOK, authEntity)

}
