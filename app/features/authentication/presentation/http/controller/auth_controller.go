package authcontroller

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	loginusecase "crosscheck-golang/app/features/authentication/domain/usecase/login"
	registrationusecase "crosscheck-golang/app/features/authentication/domain/usecase/registration"
	"crosscheck-golang/app/utils/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	registraion registrationusecase.RegistrationUsecase
	login       loginusecase.LoginUsecase
}

func New(registraion registrationusecase.RegistrationUsecase, login loginusecase.LoginUsecase) *AuthController {
	return &AuthController{
		registraion: registraion,
		login:       login,
	}
}

func (controller *AuthController) Registration(c echo.Context) error {
	param := new(param.RegistrationParam)

	log.Printf("\nRegistration controller: binding request parameters")
	if err := c.Bind(param); err != nil {
		log.Printf("\nRegistration controller: error binding! -> %+v", err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  response.ResponseStatusError,
			Message: exception.InternalServerError,
		})
	}

	log.Printf("\nRegistration controller: validating request parameters")
	if err := c.Validate(param); err != nil {
		log.Printf("\nRegistration controller: error validation! -> %+v", err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  response.ResponseStatusError,
			Message: exception.BadRequest,
		})
	}

	log.Printf("\nRegistration controller: executing registration usecase")
	authEntity, err := controller.registraion.Call(*param)

	if err != nil {
		log.Printf("\nRegistration controller: error usecase! -> %+v", err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  response.ResponseStatusError,
			Message: exception.BadRequest,
		})
	}

	log.Println("\nRegistration controller: Completed!")
	return c.JSON(http.StatusOK, response.Response{
		Status:  response.ResponseStatusSuccess,
		Message: response.ResponseMessageSuccess,
		Data:    authEntity,
	})

}

func (controller *AuthController) Login(c echo.Context) error {
	param := new(param.LoginParam)

	log.Printf("\nLogin controller: binding request parameters")
	if err := c.Bind(param); err != nil {
		log.Printf("\nLogin controller: error binding! -> %+v", err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Message: exception.InternalServerError,
		})
	}

	log.Printf("\nLogin controller: validating request parameters")
	if err := c.Validate(param); err != nil {
		log.Printf("\nLogin controller: error validation! -> %+v", err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Message: exception.BadRequest,
		})
	}

	log.Printf("\nLogin controller: executing registration usecase with params: %+v", param)
	authEntity, err := controller.login.Call(*param)

	if err != nil {
		log.Printf("\nLogin controller: error usecase! -> %+v", err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Message: exception.BadRequest,
		})
	}

	log.Println("\nLogin controller: Completed!")
	return c.JSON(http.StatusOK, response.Response{
		Message: response.RequestSuccess,
		Data:    authEntity,
	})

}
