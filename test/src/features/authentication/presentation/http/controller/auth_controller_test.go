package controller_test

import (
	"errors"
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authcontroller "crosscheck-golang/app/features/authentication/presentation/http/controller"
	mock "crosscheck-golang/test/mocks"
)

var _ = Describe("AuthController", func() {
	Describe("Registration", func() {
		var mockParam *param.RegistrationParam
		var mockParamEmpty *param.RegistrationParam
		var mockAuthEntity *entity.AuthEntity
		var mockContext *mock.MockContext
		var mockRegistrationUsecase *mock.MockRegistrationUsecase

		BeforeEach(func() {
			ctrl := gomock.NewController(GinkgoT())
			defer ctrl.Finish()

			mockContext = mock.NewMockContext(ctrl)
			mockRegistrationUsecase = mock.NewMockRegistrationUsecase(ctrl)

			mockParam = &param.RegistrationParam{
				Name:            "rizqyfahmi",
				Email:           "rizqyfahmi@email.com",
				Password:        "HelloPassword",
				ConfirmPassword: "HelloConfirmPassword",
			}

			mockParamEmpty = &param.RegistrationParam{}

			mockAuthEntity = &entity.AuthEntity{
				AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAifQ._aG0ukzancZqhL1wvBTJh8G8d3Det5n0WKcPo5C0DCY",
				RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAifQ._aG0ukzancZqhL1wvBTJh8G8d3Det5n0WKcPo5C0DCY",
			}
		})

		Context("When request successfully proccessed ", func() {
			It("returns http.ok", func() {
				mockContext.EXPECT().Bind(mockParamEmpty).SetArg(0, *mockParam).Return(nil)
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(mockAuthEntity, nil)
				mockContext.EXPECT().JSON(http.StatusOK, mockAuthEntity).Return(nil)

				app := authcontroller.New(mockRegistrationUsecase)
				err := app.Registration(mockContext)

				Expect(err).Should(Succeed())
			})
		})

		Context("When request fails to bind", func() {
			It("returns error", func() {
				mockContext.EXPECT().Bind(mockParamEmpty).Return(errors.New(exception.InternalServerError))

				app := authcontroller.New(mockRegistrationUsecase)
				err := app.Registration(mockContext)

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(exception.InternalServerError))
			})
		})

		Context("When request fails to be proccessed by usecase", func() {
			It("returns error", func() {
				mockException := &exception.Exception{
					Message: exception.ErrorDatabase,
					Causes:  gomock.Any().String(),
				}
				mockContext.EXPECT().Bind(mockParamEmpty).SetArg(0, *mockParam).Return(nil)
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(nil, mockException)

				app := authcontroller.New(mockRegistrationUsecase)
				err := app.Registration(mockContext)

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(exception.BadRequest))
			})
		})
	})
})
