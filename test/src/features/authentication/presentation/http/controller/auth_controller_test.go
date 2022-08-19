package controller_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/param"
	"crosscheck-golang/app/features/authentication/domain/entity"
	authcontroller "crosscheck-golang/app/features/authentication/presentation/http/controller"
	validatorutil "crosscheck-golang/app/utils/validator"
	mock "crosscheck-golang/test/mocks"
)

var _ = Describe("AuthController", func() {
	Describe("Registration", func() {
		var mockParam *param.RegistrationParam
		var mockAuthEntity *entity.AuthEntity
		var mockRegistrationUsecase *mock.MockRegistrationUsecase
		var e *echo.Echo

		BeforeEach(func() {
			ctrl := gomock.NewController(GinkgoT())
			defer ctrl.Finish()

			mockRegistrationUsecase = mock.NewMockRegistrationUsecase(ctrl)

			mockParam = &param.RegistrationParam{
				Name:            "Rizqy Fahmi",
				Email:           "rizqyfahmi@email.com",
				Password:        "HelloPassword",
				ConfirmPassword: "HelloPassword",
			}

			mockAuthEntity = &entity.AuthEntity{
				AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAifQ._aG0ukzancZqhL1wvBTJh8G8d3Det5n0WKcPo5C0DCY",
				RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAifQ._aG0ukzancZqhL1wvBTJh8G8d3Det5n0WKcPo5C0DCY",
			}

			e = echo.New()
			e.Validator = validatorutil.New(validator.New())
		})

		Context("When there is a valid request", func() {
			It("returns success", func() {
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(mockAuthEntity, nil).Times(1)
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusOK))

			})
		})

		Context("When there is invalid request that is caught by different content-type", func() {
			It("returns internal server error", func() {
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusInternalServerError))

			})
		})

		Context("When there is invalid request that is caught by insufficient required parameters", func() {
			It("returns bad request", func() {
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))

			})
		})

		Context("When there is invalid request that is caught by error password encryption in usecase", func() {
			It("returns bad request", func() {
				mockException := &exception.Exception{
					Message: exception.ErrorEncryption,
					Causes:  gomock.Any().String(),
				}
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(nil, mockException).Times(1)
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))

			})
		})

		Context("When there is invalid request that is caught by error database in usecase", func() {
			It("returns bad request", func() {
				mockException := &exception.Exception{
					Message: exception.ErrorDatabase,
					Causes:  gomock.Any().String(),
				}
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(nil, mockException).Times(1)
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))

			})
		})

		Context("When there is invalid request that is caught by error generate access token in usecase", func() {
			It("returns bad request", func() {
				mockException := &exception.Exception{
					Message: exception.ErrorAccessToken,
					Causes:  gomock.Any().String(),
				}
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(nil, mockException).Times(1)
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))

			})
		})

		Context("When there is invalid request that is caught by error generate refresh token in usecase", func() {
			It("returns bad request", func() {
				mockException := &exception.Exception{
					Message: exception.ErrorRefreshToken,
					Causes:  gomock.Any().String(),
				}
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(nil, mockException).Times(1)
				app := authcontroller.New(mockRegistrationUsecase)

				router := e.Group("/auth")
				router.POST("/registration", app.Registration)

				payload := `name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword`
				req, err := http.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				if err != nil {
					log.Fatal(err)
				}

				rec := httptest.NewRecorder()

				e.ServeHTTP(rec, req)

				res := rec.Result()
				defer res.Body.Close()

				Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))

			})
		})
	})
})
