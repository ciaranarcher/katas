package caesar_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ciaranarcher/katas/caesar"
)

func dotest1(s string, shift int, exp []string) {
	var ans = MovingShift(s, shift)
	Expect(ans).To(Equal(exp))
}
func dotest2(arr []string, shift int, exp string) {
	var ans = DemovingShift(arr, shift)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests", func() {
	shiftFactor := 13
	It("should handle basic cases Caesar shifting", func() {
		Expect(Shift("h", shiftFactor)).To(Equal("u"))
		Expect(Shift("x", shiftFactor)).To(Equal("k"))
		Expect(Shift("H", shiftFactor)).To(Equal("U"))
		Expect(Shift("X", shiftFactor)).To(Equal("K"))
		Expect(Shift("Z", shiftFactor)).To(Equal("M"))
		Expect(Shift(" ", shiftFactor)).To(Equal(" "))
		Expect(Shift("8", shiftFactor)).To(Equal("8"))
		Expect(Shift("hello, world", shiftFactor)).To(Equal("uryyb, jbeyq"))
	})

	It("should handle basic cases MovingShift", func() {
		var u = "I should have known that you would have a perfect answer for me!!!"
		var sol = []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}
		dotest1(u, 1, sol)
	})
	PIt("should handle basic cases MovingShift", func() {
		var u = "I should have known that you would have a perfect answer for me!!!"
		var sol = []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}
		dotest2(sol, 1, u)
	})
})
