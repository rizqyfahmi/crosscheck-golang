package authentication_repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Repository Suite")
}
