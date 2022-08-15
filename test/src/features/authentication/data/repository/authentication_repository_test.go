package authentication_repository_test

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
	var mockUserEntity *entity.UserEntity
	var mockAuthPersistentDataSource *mock.MockAuthPersistentDataSource
	var mockClock *mock.MockClock
	var authRepository authenticationRepository.AuthRepository

	mockNow := time.Now()

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()

		mockAuthPersistentDataSource = mock.NewMockAuthPersistentDataSource(ctrl)
		mockClock = mock.NewMockClock(ctrl)
		authRepository = authenticationRepositoryImpl.New(mockAuthPersistentDataSource, mockClock)
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
		mockUserEntity = &entity.UserEntity{
			Id:    mockNow.Format("20060102150405"),
			Name:  "rizqyfahmi",
			Email: "rizqyfahmi@email.com",
		}
	})

	Describe("Registration", func() {
		Context("When the authentication persistent data source returns error on insert registration data", func() {
			It("makes AuthenticationRepository returns error database", func() {
				mockClock.EXPECT().Now().Return(mockNow).AnyTimes()
				mockAuthPersistentDataSource.EXPECT().Insert(mockUserModel).Return(errors.New(exception.ErrorDatabase))

				result, err := authRepository.Registration(*mockParam)

				Expect(err).ShouldNot(BeNil())
				Expect(err.Message).Should(Equal(exception.ErrorDatabase))
				Expect(result).Should(BeNil())
			})
		})

		Context("When registration data successfully inserted by the authentication persistent data source", func() {
			It("makes AuthenticationRepository returns UserEntity", func() {
				mockClock.EXPECT().Now().Return(mockNow).AnyTimes()
				mockAuthPersistentDataSource.EXPECT().Insert(mockUserModel).Return(nil)

				result, err := authRepository.Registration(*mockParam)

				Expect(err).Should(BeNil())
				Expect(result.Id).Should(Equal(mockUserEntity.Id))
				Expect(result.Name).Should(Equal(mockUserEntity.Name))
				Expect(result.Email).Should(Equal(mockUserEntity.Email))
			})
		})
	})
})
