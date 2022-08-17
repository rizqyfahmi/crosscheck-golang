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
	return nil, nil
}
