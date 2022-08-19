package authrepository

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
)

type AuthRepository interface {
	Registration(param param.RegistrationParam) (*entity.UserEntity, *exception.Exception)
	Login(username string) (*entity.UserLoginEntity, *exception.Exception)
}
