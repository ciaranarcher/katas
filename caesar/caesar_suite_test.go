package caesar_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCaesar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Caesar Suite")
}
