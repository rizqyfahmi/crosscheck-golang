package login_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authloginuc "crosscheck-golang/app/features/authentication/domain/usecase/login"
	mock "crosscheck-golang/test/mocks"
)

var _ = Describe("LoginUsecase", func() {
	var mockParam *param.LoginParam
	var mockEntity *entity.UserLoginEntity
	var mockAuthRepository *mock.MockAuthRepository
	var mockAccessToken *mock.MockJwtUtil
	var mockRefreshToken *mock.MockJwtUtil
	var mockHash *mock.MockHash

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()

		mockAuthRepository = mock.NewMockAuthRepository(ctrl)
		mockAccessToken = mock.NewMockJwtUtil(ctrl)
		mockRefreshToken = mock.NewMockJwtUtil(ctrl)
		mockHash = mock.NewMockHash(ctrl)

		mockParam = &param.LoginParam{
			Username: "rizqyfahmi@email.com",
			Password: "HelloPassword",
		}

		mockEntity = &entity.UserLoginEntity{
			Id:       "123",
			Password: "$2a$12$JH27mGP7yo3QsV2EnRQZd.Gfx80muA71YC2f/ZVCsC/Fgr93zFvny",
		}
	})

	Context("When AuthRepository returns UserEntity", func() {
		It("makes LoginUsecase returns AuthEntity", func() {

			anyString := gomock.Any().String()
			mockAuthRepository.EXPECT().Login(mockParam.Username).Return(mockEntity, nil)
			mockHash.EXPECT().ComparePassword(mockEntity.Password, "HelloPassword").Return(nil)
			mockAccessToken.EXPECT().GenerateToken(mockEntity.Id).Return(&anyString, nil)
			mockRefreshToken.EXPECT().GenerateToken(mockEntity.Id).Return(&anyString, nil)

			usecase := authloginuc.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).Should(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(entity).ShouldNot(BeNil())
			Expect(entity.AccessToken).ShouldNot(BeEmpty())
			Expect(entity.RefreshToken).ShouldNot(BeEmpty())
		})
	})

	Context("When AuthRepository fails to get UserLoginEntity", func() {
		It("makes LoginUsecase returns ErrorDatabase", func() {
			mockException := &exception.Exception{
				Message: exception.ErrorDatabase,
				Causes:  gomock.Any().String(),
			}

			mockAuthRepository.EXPECT().Login(mockParam.Username).Return(nil, mockException)

			usecase := authloginuc.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorDatabase))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When hashed password fails to compare with password parameter", func() {
		It("makes LoginUsecase returns ErrorEncryption", func() {
			mockAuthRepository.EXPECT().Login(mockParam.Username).Return(mockEntity, nil)
			mockHash.EXPECT().ComparePassword(mockEntity.Password, "HelloPassword").Return(errors.New(gomock.Any().String()))

			usecase := authloginuc.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorEncryption))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When generating access token returns error", func() {
		It("makes LoginUsecase returns ErrorAccessToken", func() {
			anyString := gomock.Any().String()
			mockAuthRepository.EXPECT().Login(mockParam.Username).Return(mockEntity, nil)
			mockHash.EXPECT().ComparePassword(mockEntity.Password, "HelloPassword").Return(nil)
			mockAccessToken.EXPECT().GenerateToken(mockEntity.Id).Return(nil, errors.New(anyString))

			usecase := authloginuc.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorAccessToken))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When generating refresh token returns error", func() {
		It("makes LoginUsecase returns ErrorRefreshToken", func() {
			anyString := gomock.Any().String()
			mockAuthRepository.EXPECT().Login(mockParam.Username).Return(mockEntity, nil)
			mockHash.EXPECT().ComparePassword(mockEntity.Password, "HelloPassword").Return(nil)
			mockAccessToken.EXPECT().GenerateToken(mockEntity.Id).Return(&anyString, nil)
			mockRefreshToken.EXPECT().GenerateToken(mockEntity.Id).Return(nil, errors.New(anyString))

			usecase := authloginuc.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorRefreshToken))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})
})
