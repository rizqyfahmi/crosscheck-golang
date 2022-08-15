package registration_usecase_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRegistraion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registraion Suite")
}
