package auth_repository_test

import (
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/model"
	"crosscheck-golang/app/features/authentication/data/param"
	authenticationRepositoryImpl "crosscheck-golang/app/features/authentication/data/repository"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authenticationRepository "crosscheck-golang/app/features/authentication/domain/repository"
	mock "crosscheck-golang/test/mocks"
)

var _ = Describe("AuthenticationRepository", func() {

	var mockParam *param.RegistrationParam
	var mockUserModel *model.UserModel
	var mockUserLoginModel *model.UserModel
	var mockUserEntity *entity.UserEntity
	var mockUserLoginEntity *entity.UserLoginEntity
	var mockAuthPersistent *mock.MockAuthPersistent
	var mockClock *mock.MockClock
	var authRepository authenticationRepository.AuthRepository

	mockNow := time.Now()

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()

		mockAuthPersistent = mock.NewMockAuthPersistent(ctrl)
		mockClock = mock.NewMockClock(ctrl)
		authRepository = authenticationRepositoryImpl.New(mockAuthPersistent, mockClock)
		mockParam = &param.RegistrationParam{
			Name:            "rizqyfahmi",
			Email:           "rizqyfahmi@email.com",
			Password:        "HelloPassword",
			ConfirmPassword: "HelloPassword",
		}
		mockUserModel = &model.UserModel{
			Id:        mockNow.Format("20060102150405"),
			Name:      "rizqyfahmi",
			Email:     "rizqyfahmi@email.com",
			Password:  "HelloPassword",
			CreatedAt: mockNow,
			UpdatedAt: mockNow,
		}
		mockUserLoginModel = &model.UserModel{
			Id:        mockNow.Format("20060102150405"),
			Name:      "rizqyfahmi",
			Email:     "rizqyfahmi@email.com",
			Password:  "$2a$12$ZFHhRoj2.D2Mq1e9GOQRjuTplhhtKbzvhJVLLmGLeHnzkGS89wO4S",
			CreatedAt: mockNow,
			UpdatedAt: mockNow,
		}
		mockUserEntity = &entity.UserEntity{
			Id:    mockNow.Format("20060102150405"),
			Name:  "rizqyfahmi",
			Email: "rizqyfahmi@email.com",
		}
		mockUserLoginEntity = &entity.UserLoginEntity{
			Id:       mockNow.Format("20060102150405"),
			Password: "$2a$12$ZFHhRoj2.D2Mq1e9GOQRjuTplhhtKbzvhJVLLmGLeHnzkGS89wO4S",
		}
	})

	Context("Registration", func() {
		Describe("name, email, and password as the parameters", func() {
			When("Authentication repository calls insert in authentication persistent data source", func() {
				It("returns an exception that is called errorDatabase", func() {
					mockClock.EXPECT().Now().Return(mockNow).AnyTimes()
					mockAuthPersistent.EXPECT().Insert(mockUserModel).Return(errors.New(exception.ErrorDatabase))

					result, err := authRepository.Registration(*mockParam)

					Expect(err).ShouldNot(BeNil())
					Expect(err.Message).Should(Equal(exception.ErrorDatabase))
					Expect(result).Should(BeNil())
				})

				It("returns a UserEntity", func() {
					mockClock.EXPECT().Now().Return(mockNow).AnyTimes()
					mockAuthPersistent.EXPECT().Insert(mockUserModel).Return(nil)

					result, err := authRepository.Registration(*mockParam)

					Expect(err).Should(BeNil())
					Expect(result.Id).Should(Equal(mockUserEntity.Id))
					Expect(result.Name).Should(Equal(mockUserEntity.Name))
					Expect(result.Email).Should(Equal(mockUserEntity.Email))
				})
			})
		})
	})

	Context("Login", func() {
		Describe("Username as the parameters", func() {
			When("Authentication repository calls GetByUsername in authentication persistent data source", func() {
				It("returns an exception that is called errorDatabase", func() {
					mockAuthPersistent.EXPECT().GetByUsername(&mockParam.Email).Return(nil, errors.New(exception.ErrorDatabase))

					result, err := authRepository.Login(mockParam.Email)

					Expect(err).ShouldNot(BeNil())
					Expect(err.Message).Should(Equal(exception.ErrorDatabase))
					Expect(result).Should(BeNil())
				})

				It("returns a UserLoginEntity", func() {
					mockAuthPersistent.EXPECT().GetByUsername(&mockParam.Email).Return(mockUserLoginModel, nil)

					result, err := authRepository.Login(mockParam.Email)

					Expect(err).Should(BeNil())
					Expect(result.Id).Should(Equal(mockUserLoginEntity.Id))
					Expect(result.Password).Should(Equal(mockUserLoginEntity.Password))
				})
			})
		})
	})
})
