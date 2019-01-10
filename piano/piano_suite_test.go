package piano_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPiano(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Piano Suite")
}
