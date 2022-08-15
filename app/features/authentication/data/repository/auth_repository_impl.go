package authentication_repository

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/model"
	"crosscheck-golang/app/features/authentication/data/param"
	authenticationPersistent "crosscheck-golang/app/features/authentication/data/source"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authenticationRepository "crosscheck-golang/app/features/authentication/domain/repository"
	"crosscheck-golang/app/utils/clock"
)

type AuthRepositoryImpl struct {
	authLocalData authenticationPersistent.AuthPersistentDataSource
	clock         clock.Clock
}

func New(authLocalData authenticationPersistent.AuthPersistentDataSource, clock clock.Clock) authenticationRepository.AuthRepository {
	return &AuthRepositoryImpl{
		authLocalData: authLocalData,
		clock:         clock,
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

	if err := repo.authLocalData.Insert(&userModel); err != nil {
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
