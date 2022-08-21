package authusecase

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authrepository "crosscheck-golang/app/features/authentication/domain/repository"
	"crosscheck-golang/app/utils/hash"
	"crosscheck-golang/app/utils/jwt"
	"log"
)

type LoginUsecase interface {
	Call(param param.LoginParam) (*entity.AuthEntity, *exception.Exception)
}

type LoginUsecaseImpl struct {
	authRepository   authrepository.AuthRepository
	accessTokenUtil  jwt.JwtUtil
	refreshTokenUtil jwt.JwtUtil
	hash             hash.Hash
}

func New(authRepository authrepository.AuthRepository, accessTokenUtil jwt.JwtUtil, refreshTokenUtil jwt.JwtUtil, hash hash.Hash) LoginUsecase {
	return &LoginUsecaseImpl{
		authRepository,
		accessTokenUtil,
		refreshTokenUtil,
		hash,
	}
}

func (usecase *LoginUsecaseImpl) Call(param param.LoginParam) (*entity.AuthEntity, *exception.Exception) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	log.Println("\nLogin usecase: executing login repository")
	userLogin, errException := usecase.authRepository.Login(param.Username)
	if errException != nil {
		log.Println("Login usecase: executing login repository returns error!")
		return nil, errException
	}

	log.Println("Login usecase: comparing password")
	err := usecase.hash.ComparePassword(userLogin.Password, param.Password)
	if err != nil {
		log.Println("Login usecase: comparing password returns error!")
		return nil, &exception.Exception{
			Message: exception.ErrorEncryption,
			Causes:  err.Error(),
		}
	}

	var accessTokenResult = make(chan *string)
	var accessTokenErr = make(chan error)

	var refreshTokenResult = make(chan *string)
	var refreshTokenErr = make(chan error)

	defer func() {
		close(accessTokenErr)
		close(accessTokenResult)
		close(refreshTokenErr)
		close(refreshTokenResult)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()

		log.Println("Login usecase: running goroutine for generate access token")
		localResult, localError := usecase.accessTokenUtil.GenerateToken(userLogin.Id)

		accessTokenErr <- localError
		accessTokenResult <- localResult

		log.Println("Login usecase: access token successfully generated")

	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()

		log.Println("Login usecase: running goroutine for generate refresh token")
		localResult, localError := usecase.refreshTokenUtil.GenerateToken(userLogin.Id)

		refreshTokenErr <- localError
		refreshTokenResult <- localResult

		log.Println("Login usecase: refresh token successfully generated")
	}()

	err = <-accessTokenErr
	if err != nil {
		log.Println("Login usecase: generating access token returns error!")
		return nil, &exception.Exception{
			Message: exception.ErrorAccessToken,
			Causes:  err.Error(),
		}
	}

	err = <-refreshTokenErr
	if err != nil {
		log.Println("Login usecase: generating refresh token returns error!")
		return nil, &exception.Exception{
			Message: exception.ErrorRefreshToken,
			Causes:  err.Error(),
		}
	}

	accessToken := <-accessTokenResult
	refreshToken := <-refreshTokenResult

	log.Println("Login usecase: completed!")
	return &entity.AuthEntity{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}
