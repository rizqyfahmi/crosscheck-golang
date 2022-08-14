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
	user, errException := usecase.authRepository.Registration(param)
	if errException != nil {
		return nil, errException
	}

	accessToken, err := usecase.accessTokenUtil.GenerateToken(user.Id)

	if err != nil {
		return nil, &exception.Exception{
			Message: exception.ErrorAccessToken,
			Causes:  err.Error(),
		}
	}

	refeshToken, err := usecase.refreshTokenUtil.GenerateToken(user.Id)

	if err != nil {
		return nil, &exception.Exception{
			Message: exception.ErrorRefreshToken,
			Causes:  err.Error(),
		}
	}

	return &entity.AuthEntity{
		AccessToken:  *accessToken,
		RefreshToken: *refeshToken,
	}, nil
}
