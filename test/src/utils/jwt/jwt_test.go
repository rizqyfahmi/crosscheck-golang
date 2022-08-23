package jwt_test

import (
	"log"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	JwtUtil "crosscheck-golang/app/utils/jwt"
	"crosscheck-golang/config"
)

var _ = Describe("JWT", func() {
	Describe("Generate Token", func() {
		Context("When SecretKey parameter on initial is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.GenerateToken("UserID")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("SecretKey is required"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When ExpiresAt parameter on initial is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 0,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.GenerateToken("UserID")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("ExpiresAt must be greater than 0"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When UserID parameter on GenerateToken is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.GenerateToken("")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("UserID is required"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When all required parameters is not empty", func() {
			It("returns success", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.GenerateToken("UserID")

				Expect(err).Should(Succeed())
				Expect(result).ShouldNot(BeNil())
			})
		})
	})

	Describe("Validate Token", func() {
		Context("When SecretKey parameter on initial is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.ValidateToken("Token")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("SecretKey is required"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When ExpiresAt parameter on initial is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 0,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.ValidateToken("Token")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("ExpiresAt must be greater than 0"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When token parameter on ValidateToken is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.ValidateToken("")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("token is required"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When token parameter on ValidateToken is empty", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.ValidateToken("")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("token is required"))
				Expect(result).Should(BeNil())
			})
		})

		Context("When value of token parameter is invalid", func() {
			It("returns failed", func() {
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 10 * time.Minute,
				}
				jwt := JwtUtil.New[JwtUtil.JwtUtil](config)
				result, err := jwt.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

				Expect(err).Should(HaveOccurred())
				Expect(result).ShouldNot(BeNil())
			})
		})

		Context("When all required parameters is not empty", func() {
			var token string
			var jwt JwtUtil.JwtUtil

			BeforeEach(func() {
				log.Println("Hello")
				config := config.TokenConfig{
					Secret:  "SecretKey",
					Expires: 30 * time.Minute,
				}
				jwt = JwtUtil.New[JwtUtil.JwtUtil](config)
				tempToken, err := jwt.GenerateToken("UserID")

				token = *tempToken

				Expect(err).Should(Succeed())
			})

			It("returns success", func() {
				result, err := jwt.ValidateToken(token)
				Expect(err).Should(Succeed())
				Expect(result).ShouldNot(BeNil())
			})
		})
	})
})
