package auth_persistent_data_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthPersistentData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Persistent Data Suite")
}
