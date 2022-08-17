package authusecase

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authrepository "crosscheck-golang/app/features/authentication/domain/repository"
	"crosscheck-golang/app/utils/hash"
	"crosscheck-golang/app/utils/jwt"
)

type LoginUsecase struct {
	authRepository   authrepository.AuthRepository
	accessTokenUtil  jwt.JwtUtil
	refreshTokenUtil jwt.JwtUtil
	hash             hash.Hash
}

func New(authRepository authrepository.AuthRepository, accessTokenUtil jwt.JwtUtil, refreshTokenUtil jwt.JwtUtil, hash hash.Hash) *LoginUsecase {
	return &LoginUsecase{
		authRepository,
		accessTokenUtil,
		refreshTokenUtil,
		hash,
	}
}

func (usecase *LoginUsecase) Call(param param.LoginParam) (*entity.AuthEntity, *exception.Exception) {
	userLogin, errException := usecase.authRepository.Login(param.Username)
	if errException != nil {
		return nil, errException
	}

	err := usecase.hash.ComparePassword(userLogin.Password, param.Password)
	if err != nil {
		return nil, &exception.Exception{
			Message: exception.ErrorEncryption,
			Causes:  err.Error(),
		}
	}

	accessToken, err := usecase.accessTokenUtil.GenerateToken(userLogin.Id)

	if err != nil {
		return nil, &exception.Exception{
			Message: exception.ErrorAccessToken,
			Causes:  err.Error(),
		}
	}

	refeshToken, err := usecase.refreshTokenUtil.GenerateToken(userLogin.Id)

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
