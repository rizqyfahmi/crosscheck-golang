package authentication_repository

import (
	"crosscheck-golang/app/exception"
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
	return nil, nil
}
