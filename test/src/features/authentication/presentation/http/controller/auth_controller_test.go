package controller_test

import (
	"encoding/json"
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
	"crosscheck-golang/app/utils/response"
	validatorutil "crosscheck-golang/app/utils/validator"
	mock "crosscheck-golang/test/mocks"
)

var _ = Describe("AuthController", func() {
	var mockParam *param.RegistrationParam
	var mockLoginParam *param.LoginParam
	var mockAuthEntity *entity.AuthEntity
	var mockRegistrationUsecase *mock.MockRegistrationUsecase
	var mockLoginUsecase *mock.MockLoginUsecase
	var e *echo.Echo
	var app *authcontroller.AuthController

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()

		mockRegistrationUsecase = mock.NewMockRegistrationUsecase(ctrl)
		mockLoginUsecase = mock.NewMockLoginUsecase(ctrl)

		mockParam = &param.RegistrationParam{
			Name:            "Rizqy Fahmi",
			Email:           "rizqyfahmi@email.com",
			Password:        "HelloPassword",
			ConfirmPassword: "HelloPassword",
		}

		mockLoginParam = &param.LoginParam{
			Username: "rizqyfahmi@email.com",
			Password: "HelloPassword",
		}

		mockAuthEntity = &entity.AuthEntity{
			AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAifQ._aG0ukzancZqhL1wvBTJh8G8d3Det5n0WKcPo5C0DCY",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAifQ._aG0ukzancZqhL1wvBTJh8G8d3Det5n0WKcPo5C0DCY",
		}

		e = echo.New()
		e.Validator = validatorutil.New(validator.New())

		app = authcontroller.New(mockRegistrationUsecase, mockLoginUsecase)
		router := e.Group("/auth")
		router.POST("/registration", app.Registration)
		router.POST("/login", app.Login)
	})

	Describe("Registration", func() {
		Context("When there is a valid request", func() {
			It("returns success", func() {
				mockRegistrationUsecase.EXPECT().Call(*mockParam).Return(mockAuthEntity, nil).Times(1)

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

	Context("Login", func() {
		Describe("Username and password as the request parameters", func() {
			When("The request parameter can't be binded caught by different content-type", func() {
				It("returns internal server error", func() {
					payload := `username=rizqyfahmi@email.com&password=HelloPassword`
					req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(payload))
					req.Header.Set("Content-Type", "application/json")

					if err != nil {
						log.Fatal(err)
					}

					rec := httptest.NewRecorder()

					e.ServeHTTP(rec, req)

					res := rec.Result()
					defer res.Body.Close()

					result := response.Response{}
					_ = json.Unmarshal(rec.Body.Bytes(), &result)

					Expect(res).Should(HaveHTTPStatus(http.StatusInternalServerError))
					Expect(result.Message).Should(Equal(exception.InternalServerError))
				})
			})
			When("The request parameter can't be binded caught by error validation", func() {
				It("returns bad request", func() {
					payload := `username=rizqyfahmi@email.com&password=`
					req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(payload))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

					if err != nil {
						log.Fatal(err)
					}

					rec := httptest.NewRecorder()

					e.ServeHTTP(rec, req)

					res := rec.Result()
					defer res.Body.Close()

					result := response.Response{}
					_ = json.Unmarshal(rec.Body.Bytes(), &result)

					Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))
					Expect(result.Message).Should(Equal(exception.BadRequest))
				})
			})
			When("The usecase fails to get data from database", func() {
				It("returns bad request", func() {
					mockException := &exception.Exception{
						Message: exception.ErrorDatabase,
						Causes:  gomock.Any().String(),
					}
					mockLoginUsecase.EXPECT().Call(*mockLoginParam).Return(nil, mockException).Times(1)

					payload := `username=rizqyfahmi@email.com&password=HelloPassword`
					req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(payload))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

					if err != nil {
						log.Fatal(err)
					}

					rec := httptest.NewRecorder()

					e.ServeHTTP(rec, req)

					res := rec.Result()
					defer res.Body.Close()

					result := response.Response{}
					_ = json.Unmarshal(rec.Body.Bytes(), &result)

					Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))
					Expect(result.Message).Should(Equal(exception.BadRequest))
				})
			})
			When("The usecase fails to generate access token", func() {
				It("returns bad request", func() {
					mockException := &exception.Exception{
						Message: exception.ErrorAccessToken,
						Causes:  gomock.Any().String(),
					}
					mockLoginUsecase.EXPECT().Call(*mockLoginParam).Return(nil, mockException).Times(1)

					payload := `username=rizqyfahmi@email.com&password=HelloPassword`
					req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(payload))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

					if err != nil {
						log.Fatal(err)
					}

					rec := httptest.NewRecorder()

					e.ServeHTTP(rec, req)

					res := rec.Result()
					defer res.Body.Close()

					result := response.Response{}
					_ = json.Unmarshal(rec.Body.Bytes(), &result)

					Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))
					Expect(result.Message).Should(Equal(exception.BadRequest))
				})
			})
			When("The usecase fails to generate refresh token", func() {
				It("returns bad request", func() {
					mockException := &exception.Exception{
						Message: exception.ErrorRefreshToken,
						Causes:  gomock.Any().String(),
					}
					mockLoginUsecase.EXPECT().Call(*mockLoginParam).Return(nil, mockException).Times(1)

					payload := `username=rizqyfahmi@email.com&password=HelloPassword`
					req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(payload))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

					if err != nil {
						log.Fatal(err)
					}

					rec := httptest.NewRecorder()

					e.ServeHTTP(rec, req)

					res := rec.Result()
					defer res.Body.Close()

					result := response.Response{}
					_ = json.Unmarshal(rec.Body.Bytes(), &result)

					Expect(res).Should(HaveHTTPStatus(http.StatusBadRequest))
					Expect(result.Message).Should(Equal(exception.BadRequest))
				})
			})
			When("The login controller successfully process the request", func() {
				It("returns AuthEntity", func() {
					mockLoginUsecase.EXPECT().Call(*mockLoginParam).Return(mockAuthEntity, nil).Times(1)

					payload := `username=rizqyfahmi@email.com&password=HelloPassword`
					req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(payload))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

					if err != nil {
						log.Fatal(err)
					}

					rec := httptest.NewRecorder()

					e.ServeHTTP(rec, req)

					res := rec.Result()
					defer res.Body.Close()

					result := response.Response{}
					_ = json.Unmarshal(rec.Body.Bytes(), &result)

					data := entity.AuthEntity{}
					dataJSON, _ := json.Marshal(result.Data)
					_ = json.Unmarshal(dataJSON, &data)

					Expect(rec).Should(HaveHTTPStatus(http.StatusOK))
					Expect(data.AccessToken).Should(Equal(mockAuthEntity.AccessToken))
					Expect(data.RefreshToken).Should(Equal(mockAuthEntity.RefreshToken))
				})
			})
		})
	})
})
