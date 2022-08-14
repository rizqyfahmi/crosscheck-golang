package utils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"

	"crosscheck-golang/app/utils"
)

var _ = Describe("Bcrypt", func() {
	Describe("Hash Password", func() {
		Context("When the result is compared with different password", func() {
			It("returns failure", func() {
				hashPassword := utils.HashPassword("HelloPassword")
				err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte("OllaPassword"))

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("crypto/bcrypt: hashedPassword is not the hash of the given password"))
			})
		})

		Context("When the result is compared with similar password", func() {
			It("returns success", func() {
				hashPassword := utils.HashPassword("HelloPassword")
				err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte("HelloPassword"))

				Expect(err).Should(Succeed())
			})
		})

		Context("When the password parameter is already encrypted", func() {
			It("returns failure", func() {
				hashPassword := utils.HashPassword("$2a$12$kSmt/kAF0Yf0egtWzvWQR.XOcpy0QkG7qe5BWKfCua.nUw3fqguSS")
				err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte("OllaPassword"))

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("crypto/bcrypt: hashedPassword is not the hash of the given password"))
			})
		})
	})

	Describe("Compare password", func() {
		var result []byte

		BeforeEach(func() {
			pw := []byte("HelloPassword")
			tempResult, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

			result = tempResult

			Expect(err).Should(Succeed())
		})

		Context("When the password parameters are not equal", func() {
			It("returns failure", func() {
				err := utils.ComparePassword(string(result), "OllaPassword")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError("crypto/bcrypt: hashedPassword is not the hash of the given password"))
			})
		})

		Context("When the password parameters are equal", func() {
			It("returns success", func() {
				err := utils.ComparePassword(string(result), "HelloPassword")

				Expect(err).Should(Succeed())
			})
		})
	})
})
