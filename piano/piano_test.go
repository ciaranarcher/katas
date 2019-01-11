package piano_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ciaranarcher/katas/piano"
)

var _ = Describe("Example tests", func() {
	It("should test that the solution returns the correct value", func() {

		// Deal with the edges
		Expect(BlackOrWhiteKey(1)).To(Equal("white"))
		Expect(BlackOrWhiteKey(2)).To(Equal("black"))
		Expect(BlackOrWhiteKey(3)).To(Equal("white"))
		Expect(BlackOrWhiteKey(88)).To(Equal("white"))

		Expect(BlackOrWhiteKey(4)).To(Equal("white"))
		Expect(BlackOrWhiteKey(5)).To(Equal("black"))
		Expect(BlackOrWhiteKey(6)).To(Equal("white"))
		Expect(BlackOrWhiteKey(7)).To(Equal("black"))
		Expect(BlackOrWhiteKey(8)).To(Equal("white"))
		Expect(BlackOrWhiteKey(9)).To(Equal("white"))
		Expect(BlackOrWhiteKey(10)).To(Equal("black"))
		Expect(BlackOrWhiteKey(11)).To(Equal("white"))
		Expect(BlackOrWhiteKey(12)).To(Equal("black"))
		Expect(BlackOrWhiteKey(13)).To(Equal("white"))
		Expect(BlackOrWhiteKey(14)).To(Equal("black"))
		Expect(BlackOrWhiteKey(15)).To(Equal("white"))
		Expect(BlackOrWhiteKey(16)).To(Equal("white"))

		Expect(BlackOrWhiteKey(5)).To(Equal("black"))
		Expect(BlackOrWhiteKey(12)).To(Equal("black"))
		Expect(BlackOrWhiteKey(42)).To(Equal("white"))
		Expect(BlackOrWhiteKey(88)).To(Equal("white"))
		Expect(BlackOrWhiteKey(89)).To(Equal("white"))
		Expect(BlackOrWhiteKey(92)).To(Equal("white"))
		Expect(BlackOrWhiteKey(100)).To(Equal("black"))
		Expect(BlackOrWhiteKey(111)).To(Equal("white"))
		Expect(BlackOrWhiteKey(200)).To(Equal("black"))
		Expect(BlackOrWhiteKey(2017)).To(Equal("white"))
	})
})
