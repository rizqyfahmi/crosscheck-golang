package authusecase

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authenticationRepository "crosscheck-golang/app/features/authentication/domain/repository"
	"crosscheck-golang/app/utils/hash"
	jwtUtil "crosscheck-golang/app/utils/jwt"
)

type RegistrationUsecase struct {
	authRepository   authenticationRepository.AuthRepository
	accessTokenUtil  jwtUtil.JwtUtil
	refreshTokenUtil jwtUtil.JwtUtil
	hash             hash.Hash
}

func New(authRepository authenticationRepository.AuthRepository, accessTokenUtil jwtUtil.JwtUtil, refreshTokenUtil jwtUtil.JwtUtil, hash hash.Hash) *RegistrationUsecase {
	return &RegistrationUsecase{
		authRepository,
		accessTokenUtil,
		refreshTokenUtil,
		hash,
	}
}

func (usecase *RegistrationUsecase) Call(param param.RegistrationParam) (*entity.AuthEntity, *exception.Exception) {

	hashedPassword, err := usecase.hash.HashPassword(param.Password)

	if err != nil {
		return nil, &exception.Exception{
			Message: exception.ErrorEncryption,
			Causes:  err.Error(),
		}
	}

	param.Password = *hashedPassword

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
