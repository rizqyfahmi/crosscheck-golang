package usecase

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	"crosscheck-golang/app/features/authentication/domain/repository"
	jwtUtil "crosscheck-golang/app/utils/jwt"
)

type RegistrationUsecase struct {
	authRepository   repository.AuthRepository
	accessTokenUtil  jwtUtil.JwtUtil
	refreshTokenUtil jwtUtil.JwtUtil
}

func New(authRepository repository.AuthRepository, accessTokenUtil jwtUtil.JwtUtil, refreshTokenUtil jwtUtil.JwtUtil) *RegistrationUsecase {
	return &RegistrationUsecase{
		authRepository,
		accessTokenUtil,
		refreshTokenUtil,
	}
}

func (usecase *RegistrationUsecase) Call(param param.RegistrationParam) (*entity.AuthEntity, *exception.Exception) {
	return nil, nil
}
