package registration_usecase_test

import (
	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	RegistationUsecase "crosscheck-golang/app/features/authentication/domain/usecase"
	mock "crosscheck-golang/test/mocks"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationUsecase", func() {
	var mockParam *param.RegistrationParam
	var mockEntity *entity.UserEntity
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

		mockParam = &param.RegistrationParam{
			Name:            "rizqyfahmi",
			Email:           "rizqyfahmi@email.com",
			Password:        "HelloPassword",
			ConfirmPassword: "HelloConfirmPassword",
		}

		mockEntity = &entity.UserEntity{
			Id:    "123",
			Name:  "rizqyfahmi",
			Email: "rizqyfahmi@email.com",
		}

	})

	Context("When AuthRepository returns UserEntity", func() {
		It("makes RegistrationUsecase returns AuthEntity", func() {

			anyString := gomock.Any().String()
			mockHash.EXPECT().HashPassword(mockParam.Password).Return(&anyString, nil)
			mockAuthRepository.EXPECT().Registration(gomock.Any()).Return(mockEntity, nil)
			mockAccessToken.EXPECT().GenerateToken(mockEntity.Id).Return(&anyString, nil)
			mockRefreshToken.EXPECT().GenerateToken(mockEntity.Id).Return(&anyString, nil)

			usecase := RegistationUsecase.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).Should(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(entity).ShouldNot(BeNil())
			Expect(entity.AccessToken).ShouldNot(BeEmpty())
			Expect(entity.RefreshToken).ShouldNot(BeEmpty())
		})
	})

	Context("When password fails to hash", func() {
		It("makes RegistrationUsecase returns ErrorEncryption", func() {
			mockHash.EXPECT().HashPassword(mockParam.Password).Return(nil, errors.New(gomock.Any().String()))

			usecase := RegistationUsecase.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorEncryption))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When AuthRepository insert data returns error", func() {
		It("makes RegistrationUsecase returns ErrorDatabase", func() {
			mockException := &exception.Exception{
				Message: exception.ErrorDatabase,
				Causes:  gomock.Any().String(),
			}

			anyString := gomock.Any().String()
			mockHash.EXPECT().HashPassword(mockParam.Password).Return(&anyString, nil)
			mockAuthRepository.EXPECT().Registration(gomock.Any()).Return(nil, mockException)

			usecase := RegistationUsecase.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorDatabase))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When AuthRepository is mapping the inserted data into UserEntity returns error", func() {
		It("makes RegistrationUsecase returns ErrorStructMapping", func() {
			mockException := &exception.Exception{
				Message: exception.ErrorStructMapping,
				Causes:  gomock.Any().String(),
			}

			anyString := gomock.Any().String()
			mockHash.EXPECT().HashPassword(mockParam.Password).Return(&anyString, nil)
			mockAuthRepository.EXPECT().Registration(gomock.Any()).Return(nil, mockException)

			usecase := RegistationUsecase.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorStructMapping))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When generating access token returns error", func() {
		It("makes RegistrationUsecase returns ErrorAccessToken", func() {
			anyString := gomock.Any().String()
			mockHash.EXPECT().HashPassword(mockParam.Password).Return(&anyString, nil)
			mockAuthRepository.EXPECT().Registration(gomock.Any()).Return(mockEntity, nil)
			mockAccessToken.EXPECT().GenerateToken(mockEntity.Id).Return(nil, errors.New(anyString))

			usecase := RegistationUsecase.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorAccessToken))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})

	Context("When generating refresh token returns error", func() {
		It("makes RegistrationUsecase returns ErrorRefreshToken", func() {
			anyString := gomock.Any().String()
			mockHash.EXPECT().HashPassword(mockParam.Password).Return(&anyString, nil)
			mockAuthRepository.EXPECT().Registration(gomock.Any()).Return(mockEntity, nil)
			mockAccessToken.EXPECT().GenerateToken(mockEntity.Id).Return(&anyString, nil)
			mockRefreshToken.EXPECT().GenerateToken(mockEntity.Id).Return(nil, errors.New(anyString))

			usecase := RegistationUsecase.New(mockAuthRepository, mockAccessToken, mockRefreshToken, mockHash)
			entity, err := usecase.Call(*mockParam)

			Expect(err).ShouldNot(BeNil()) // We don't need Succeed, because we use a custom error struct
			Expect(err.Message).Should(Equal(exception.ErrorRefreshToken))
			Expect(err.Causes).Should(Equal(gomock.Any().String()))
			Expect(entity).Should(BeNil())
		})
	})
})
