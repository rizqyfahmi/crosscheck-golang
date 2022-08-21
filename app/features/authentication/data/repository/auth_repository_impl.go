package authrepository

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/model"
	"crosscheck-golang/app/features/authentication/data/param"
	authpersistent "crosscheck-golang/app/features/authentication/data/source/persistent"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authepository "crosscheck-golang/app/features/authentication/domain/repository"
	"crosscheck-golang/app/utils/clock"
	"log"
)

type AuthRepositoryImpl struct {
	authPersistent authpersistent.AuthPersistent
	clock          clock.Clock
}

func New(authPersistent authpersistent.AuthPersistent, clock clock.Clock) authepository.AuthRepository {
	return &AuthRepositoryImpl{
		authPersistent: authPersistent,
		clock:          clock,
	}
}

func (repo *AuthRepositoryImpl) Registration(param param.RegistrationParam) (*entity.UserEntity, *exception.Exception) {

	userModel := model.UserModel{
		Id:        repo.clock.Now().Format("20060102150405"),
		Name:      param.Name,
		Email:     param.Email,
		Password:  param.Password,
		CreatedAt: repo.clock.Now(),
		UpdatedAt: repo.clock.Now(),
	}

	if err := repo.authPersistent.Insert(&userModel); err != nil {
		return nil, &exception.Exception{
			Message: exception.ErrorDatabase,
			Causes:  err.Error(),
		}
	}

	userEntity := entity.UserEntity{
		Id:    userModel.Id,
		Name:  userModel.Name,
		Email: userModel.Email,
	}

	return &userEntity, nil
}

func (repo *AuthRepositoryImpl) Login(username string) (*entity.UserLoginEntity, *exception.Exception) {

	log.Println("\nAuthentication repository: executing GetByUsername in authentication persistent data source")
	userModel, err := repo.authPersistent.GetByUsername(&username)

	if err != nil {
		log.Printf("\nAuthentication repository: error persistent data source! -> %+v", err)
		return nil, &exception.Exception{
			Message: exception.ErrorDatabase,
			Causes:  err.Error(),
		}
	}

	log.Println("\nAuthentication repository: converting user model into user login entity")
	userLoginEntity := entity.UserLoginEntity{
		Id:       userModel.Id,
		Password: userModel.Password,
	}

	log.Println("\nAuthentication repository: completed!")
	return &userLoginEntity, nil
}
